import { useParams } from "react-router";
import PostCard from "../components/post-card";
import { useQuery } from "@tanstack/react-query";
import { fetchUser, fetchUserPosts } from "../services/api";
import CreateNewPostCard from "../components/create-newpost";
import BackArrowIcon from "../assets/icons/arrow-left.svg";
import { useNavigate } from "react-router";
import LoadingEllipsis from "../components/loadingEllipsis";
import { useState } from "react";
import Modal from "../components/new-post-modal";

interface PostProps {
  id?: string;
  userId: string;
  title: string;
  body: string;
  created_at?: string;
}

function Post() {
  const { userID } = useParams();
  const navigate = useNavigate();
  const [isModalOpen, setIsModalOpen] = useState(false);

  const handleOpenModal = () => {
    setIsModalOpen(true);
  };

  const handleCloseModal = () => {
    setIsModalOpen(false); // Close the modal
  };

  const { isPending, isError, error, data } = useQuery({
    queryKey: ["userPosts", userID],
    queryFn: () => fetchUserPosts(userID ?? ""),
  });

  const user = useQuery({
    queryKey: ["userDetails", userID],
    queryFn: () => fetchUser(userID ?? ""),
  });

  const handleBackClick = () => {
    void navigate("/users");
  };

  return (
    <>
      {/* Loading State */}
      {isPending ? (
        <div className="absolute top-[321px] left-0 right-0 flex justify-center">
          <LoadingEllipsis />
        </div>
      ) : isError ? (
        <div className="absolute top-[321px] left-0 right-0 flex justify-center text-red-500">
          <p>
            {error.message || "Failed to fetch data. Please try again later"}.
          </p>{" "}
        </div>
      ) : (
        // Content if data is loaded
        <div className="flex justify-center text-[#535862] py-4 md:py-10">
          <div className="flex flex-col gap-6">
            <div className="flex flex-col gap-4 text-left">
              <button
                onClick={handleBackClick}
                className="flex gap-1 items-center text-sm font-semibold"
              >
                <img
                  className="cursor-pointer"
                  src={BackArrowIcon}
                  alt="back"
                />
                <span className="cursor-pointer">Back to Users</span>
              </button>
              <h1 className="text-3xl font-medium text-black">
                {user.data?.firstname} {user.data?.lastname}
              </h1>
              <div className="text-sm">
                <span>{user.data?.email}</span>
                <span> &bull; </span>
                <span className="font-medium">
                  {data?.length ?? 0} Post{(data?.length ?? 0) !== 1 && "s"}
                </span>
              </div>
            </div>
            <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-y-6 gap-x-[23px] w-fit">
              <CreateNewPostCard onClick={handleOpenModal} />
              {isModalOpen && (
                <Modal
                  userId={user.data?.id ?? ""}
                  handleCloseModal={handleCloseModal}
                />
              )}
              {data?.map((post: PostProps, index: number) => (
                <PostCard
                  userID={userID ?? ""}
                  id={post.id ?? ""}
                  title={post.title}
                  content={post.body}
                  key={index}
                  onEdit={() => console.log("Implement me")}
                />
              ))}
            </div>
          </div>
        </div>
      )}
    </>
  );
}

export default Post;
