import { useParams, useLocation, useNavigate } from "react-router-dom";
import PostCard from "../components/post-card";
import { useQuery, useMutation, useQueryClient } from "@tanstack/react-query";
import { fetchUser, fetchUserPosts, deletePost } from "../services/api";
import CreateNewPostCard from "../components/create-newpost";
import BackArrowIcon from "../assets/icons/arrow-left.svg";
import LoadingEllipsis from "../components/loadingEllipsis";
import { useState } from "react";
import Modal from "../components/new-post-modal";
import Toast from "../components/toast";

interface PostProps {
  id?: string;
  userId: string;
  title: string;
  body: string;
  created_at?: string;
}

function Post() {
  const { userID } = useParams();
  const location = useLocation();
  const navigate = useNavigate();
  const queryClient = useQueryClient();
  const [isModalOpen, setIsModalOpen] = useState(false);
  const [toastMessage, setToastMessage] = useState("");
  const [toastType, setToastType] = useState<"success" | "error">("success");

  const handleOpenModal = () => setIsModalOpen(true);
  const handleCloseModal = () => setIsModalOpen(false);

  const { isLoading, isError, error, data } = useQuery({
    queryKey: ["userPosts", userID],
    queryFn: () => fetchUserPosts(userID ?? ""),
  });

  const user = useQuery({
    queryKey: ["userDetails", userID],
    queryFn: () => fetchUser(userID ?? ""),
  });

  // Retrieve the previous page from location.state (default to 1)
  const previousPage = (location.state as { page?: number })?.page || 1;

  const handleBackClick = () => {
    // Navigate back to Users with the previous page in the query string.
    navigate(`/users?page=${previousPage}`);
  };

  const deleteMutation = useMutation({
    mutationFn: (postId: string) => deletePost(postId),
    onSuccess: async () => {
      await queryClient.invalidateQueries({ queryKey: ["userPosts", userID] });
      setToastMessage("Post deleted successfully!");
      setToastType("success");
    },
    onError: () => {
      setToastMessage("Failed to delete post. Please try again.");
      setToastType("error");
    },
  });

  const handleDeletePost = (postId: string) => {
    deleteMutation.mutate(postId);
  };

  return (
    <>
      {toastMessage && (
        <Toast
          message={toastMessage}
          type={toastType}
          onClose={() => setToastMessage("")}
        />
      )}

      {isLoading ? (
        <div className="absolute top-[321px] left-0 right-0 flex justify-center">
          <LoadingEllipsis />
        </div>
      ) : isError ? (
        <div className="absolute top-[321px] left-0 right-0 flex justify-center text-red-500">
          <p>
            {error.message || "Failed to fetch data. Please try again later"}.
          </p>
        </div>
      ) : (
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
                  setToastMessage={(message: string) => {
                    setToastMessage(message);
                    setToastType("success");
                  }}
                />
              )}
              {data?.map((post: PostProps, index: number) => (
                <PostCard
                  userID={userID ?? ""}
                  id={post.id ?? ""}
                  title={post.title}
                  content={post.body}
                  key={index}
                  onDelete={() => handleDeletePost(post.id ?? "")}
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
