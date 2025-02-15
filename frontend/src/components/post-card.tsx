import React from "react";
import DeleteIcon from "../assets/icons/delete.svg";

interface PostCardProps {
  userID: string;
  id: string;
  title: string;
  content: string;
  onEdit?: () => void;
  onDelete: () => void;
}

const PostCard: React.FC<PostCardProps> = ({
  title,
  content,
  onDelete,
}) => {
  const handleDelete = (e: React.MouseEvent) => {
    e.stopPropagation();
    onDelete();
  };

  return (
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
  );
};

export default PostCard;
