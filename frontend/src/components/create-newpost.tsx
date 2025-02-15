import PlusIcon from "../assets/icons/plus.svg";

const CreateNewPostCard = ({ onClick }: { onClick: () => void }) => {
  return (
    <div
      className="flex flex-col items-center justify-center gap-2 bg-white border border-dashed border-[#D5D7DA] text-[#535862] rounded-lg cursor-pointer py-24 w-[270px] h-[293px]"
      onClick={onClick}
      role="button"
    >
      <img src={PlusIcon} alt="add a new post" width={24} height={24} />
      <span className="text-sm leading-[20px] font-semibold font-semibold">
        New Post
      </span>
    </div>
  );
};


export default CreateNewPostCard;
