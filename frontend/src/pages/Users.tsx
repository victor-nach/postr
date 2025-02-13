import { useState } from "react";
import { keepPreviousData, useQuery } from "@tanstack/react-query";
import Pagination from "../components/pagination";
import { fetchUsers } from "../services/api";
import Layout from "../components/layout";
import LoadingEllipsis from "../components/loadingEllipsis";
import { useNavigate } from "react-router";

function Users() {

  const [page, setPage] = useState<number>(1);
  const navigate = useNavigate();

  const { isPending, isError, error, data } = useQuery({
    queryKey: ["users", page],
    queryFn: () => fetchUsers(page),
    placeholderData: keepPreviousData,
  });

  const handlePostClick = (userID: string) => {
     void navigate(`/users/${userID}/posts`);
  };


  return (
    <Layout title="Users">
      <div className="border rounded-lg border-[#E9EAEB] w-full mx-auto px-4 my-[24px] overflow-x-auto">
        <table className="w-full">
          <thead className="text-left text-[#535862]">
            <tr>
              <th className="text-xs leading-[18px] px-6 font-medium">
                Full Name
              </th>
              <th className="text-xs leading-[18px] px-6 py-3 font-medium">
                Email Address
              </th>
              <th className="text-xs leading-[18px] px-6 py-3 font-medium w-[392px]">
                Address
              </th>
            </tr>
          </thead>

          {isPending ? (
            <tbody className="text-left text-[#535862]">
              <tr>
                <td colSpan={3} className="px-6 py-6 text-center">
                  <LoadingEllipsis />
                </td>
              </tr>
            </tbody>
          ) : isError ? (
            <tbody className="text-left text-[#535862]">
              <tr>
                <td colSpan={3} className="px-6 py-6 text-center">
                  <p>
                    {error.message ||
                      "Failed to fetch data. Please try again later"}
                    .
                  </p>
                </td>
              </tr>
            </tbody>
          ) : (
            <tbody className="text-left text-[#535862] text-sm leading-[14px]">
              {data.users.map((user) => (
                <tr
                  key={user.id}
                  className="border-b border-[#E9EAEB] last:border-0 cursor-pointer"
                  onClick={() => handlePostClick(user.id)}
                >
                  <td className="px-6 py-[26px] w-[124px] max-w-[124px] md:w-[200px] md:max-w-[200px] truncate">
                    {user.firstname} {user.lastname}
                  </td>
                  <td className="px-6 py-[26px] w-[124px] max-w-[124px] md:w-[264px] md:max-w-[264px] truncate">
                    {user.email}
                  </td>
                  <td className="px-6 py-[26px] w-[392px] max-w-[392px] truncate">
                      {user.street}, {user.state}, {user.city}, {user.zipcode}
                  </td>
                </tr>
              ))}
            </tbody>
          )}
        </table>
      </div>
      <Pagination
        currentPage={data?.currentPage ?? 1}
        totalPages={data?.totalPages ?? 1}
        setPage={setPage}
      />
    </Layout>
  );
}

export default Users;
