import {  useMutation, useQueryClient } from "@tanstack/react-query";
import DeleteIcon from "../assets/icons/delete.svg";
import { deletePost } from "../services/api";

interface PostCardProps {
  id: string;
  title: string;
  content: string;
  onEdit?: () => void;
}

const PostCard = ({ title, content, id }: PostCardProps) => {
  const queryClient = useQueryClient();

  const mutation = useMutation({

    mutationFn: async() => {
       await deletePost(id)
    },

    onSuccess: async () => {
      await queryClient.invalidateQueries({ queryKey: ["userPosts", id] });

     console.log('data deleted')
    },

    onError: (error) => {
      console.error("Error deleting post:", error);
    },
  });

  const handleDelete = async () => {
    try {
      await mutation.mutateAsync(); // Await the mutation
      console.log("Post deleted successfully");
    } catch (error) {
      console.error("Error deleting post:", error);
    }
  };

  return (
    <div className="relative flex flex-col gap-4 p-3 sm:p-6 rounded-lg bg-white border border-[#D5D7DA] text-[#535862] text-left cursor-pointer w-[270px] h-[300px]">
      <div className="absolute top-1 right-1 flex gap-2">
        <button className="cursor-pointer"
          onClick={(e) => {
            e.stopPropagation();
            void handleDelete()
          }}
        >
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
  );
};

export default PostCard;
