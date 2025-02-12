import { useMutation, useQueryClient } from "@tanstack/react-query";
import DeleteIcon from "../assets/icons/delete.svg";
import { deletePost } from "../services/api";
import { useState } from "react";
// import { toast, ToastContainer } from "react-toastify";

interface PostCardProps {
  userID: string;
  id: string;
  title: string;
  content: string;
  onEdit?: () => void;
}

const PostCard = ({ title, content, id, userID }: PostCardProps) => {
  const queryClient = useQueryClient();
  const [toast, setToast] = useState<{
    type: "success" | "error";
    message: string;
  } | null>(null);

  const mutation = useMutation({
    mutationFn: async () => {
      return deletePost(id);
    },

    onMutate: async () => {
      await queryClient.cancelQueries({ queryKey: ["userPosts", userID] });

      // Optimistically update the cache to remove the deleted post
      const previousPosts = queryClient.getQueryData(["userPosts", userID]);

      if (previousPosts) {
        queryClient.setQueryData(
          ["userPosts", userID],
          (previousPosts: PostCardProps[]) => {
            return previousPosts.filter((post) => post.id !== id);
          }
        );
      }
      return { previousPosts };
    },

    onSuccess: () => {
      setToast({ type: "success", message: "Post deleted successfully!" });
      setTimeout(() => setToast(null), 3000);
      void queryClient.invalidateQueries({ queryKey: ["userPosts", userID] });
    },

    onError: (error) => {
      console.error("Error deleting post:", error);

      setToast({ type: "error", message: "Could not delete post" });
      setTimeout(() => setToast(null), 3000);
    },
  });

  const handleDelete = (e: React.MouseEvent) => {
    e.stopPropagation();
    mutation.mutate();
  };

  return (
    <>
      {toast && (
        <div
          className={`fixed top-5 right-5 px-4 py-2 rounded-lg shadow-lg transition-all flex items-center justify-between w-60 ${
            toast.type === "success"
              ? "bg-green-500 text-white"
              : "bg-red-500 text-white"
          }`}
        >
          <span className="text-sm font-medium">{toast.message}</span>

          <button
            className="bg-white text-green-600 hover:bg-gray-200 rounded-full w-6 h-6 flex items-center justify-center text-sm font-bold"
            onClick={() => setToast(null)}
          >
            ✕
          </button>
        </div>
      )}

      <div className="relative flex flex-col gap-4 p-3 sm:p-6 rounded-lg bg-white border border-[#D5D7DA] text-[#535862] text-left cursor-pointer w-[270px] h-[300px]">
        <div className="absolute top-1 right-1 flex gap-2">
          <button className="cursor-pointer" onClick={handleDelete}>
            <img src={DeleteIcon} alt="delete" />
          </button>
        </div>
        <h1 className="text-lg font-medium line-clamp-2 break-all pr-[21px]">
          {title}
        </h1>
        <p className="text-sm leading-normal line-clamp-[8] break-all">
          {content}
        </p>
      </div>
    </>
  );
};

export default PostCard;
