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
      <div className="border rounded-lg border-[#E9EAEB] max-w-[856px] w-full  overflow-auto my-[24px]">
        <table>
          <thead className="text-left text-[#535862]">
            <tr>
              <th className="text-xs leading-[18px] md:px-6 py-3 font-medium  sm:w-[124px] max-w-[124px] md:w-[200px]">
                Full Name
              </th>
              <th className="text-xs leading-[18px] md:px-6 py-3 font-medium  sm:w-[124px] md:w-[264px] ">
                Email Address
              </th>
              <th className="text-xs leading-[18px] md:px-6 py-3 font-medium w-[392px]">
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
                  <td className="font-medium px-6 py-[26px] sm:w-[124px] sm:max-w-[124px] md:w-[200px] md:max-w-[200px] overflow-auto whitespace-nowrap text-ellipsis">
                    {user.firstname} {user.lastname}
                  </td>
                  <td className="font-normal px-6 py-[26px] sm:w-[124px] sm:max-w-[124px] md:w-[264px] md:max-w-[264px] overflow-auto text-ellipsis whitespace-nowrap">
                    {user.email}
                  </td>
                  <td className="px-6 py-[26px] w-[392px] max-w-[392px]">
                    <p className="text-ellipsis whitespace-nowrap overflow-auto ">
                      {user.street}, {user.state}, {user.city}, {user.zipcode}
                    </p>
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
