import { useMutation, useQueryClient } from "@tanstack/react-query";
import { createPost } from "../services/api";
import { useState } from "react";
import { validatePost } from "../utils";

export default function Modal({
  handleCloseModal,
  userId,
}: {
  userId: string;
  handleCloseModal: () => void;
}) {
  const [title, setTitle] = useState<string>("");
  const [body, setBody] = useState("");
  const [validationError, setValidationError] = useState("");
  const queryClient = useQueryClient();

  const { isPending, mutate, isError } = useMutation({
    mutationFn: async () => await createPost(userId, title, body),

    onSuccess: async () => {
      await queryClient.invalidateQueries({ queryKey: ["userPosts", userId] });
      handleCloseModal();
      setTitle("");
      setBody("");
    },

    onError: (error) => {
      setTitle("");
      setBody("");
      console.error("Error deleting post:", error);
    },
  });

  const handleSubmit = () => {
    console.log('licke')
    const validationError = validatePost(title, body);

    if (validationError) {
      setValidationError(validationError);
      return;
    }
    setValidationError("");
    mutate();
   
  };

  return (
    <div className="fixed top-0 left-0 right-0 bottom-0 flex justify-center items-center bg-[#03071282]  z-50">
      <div className="bg-white p-[24px] rounded-lg w-[679px]">
        <h2 className="text-4xl tracking-[-2%] leading-[43.57px] text-[#181D27] font-medium">
          New Post
        </h2>
        <form onSubmit={handleSubmit}>
          <div className="pt-5">
            <label
              htmlFor="post-title"
              className="font-medium leading-[20px] text-lg text-[#535862]"
            >
              Post title
              <input
                onChange={(e) => setTitle(e.target.value)}
                type="text"
                name="title"
                value={title}
                placeholder="Give your post a title"
                className="border border-[#E2E8F0] rounded-lg p-2 mt-1 w-full text-sm font-normal leading-[21px] placeholder:text-sm placeholder:leading-[21px] placeholder:font-normal px-3"
              />
            </label>
          </div>
          <div className="pt-4">
            <label
              htmlFor="post-content"
              className="font-medium leading-[20px] text-lg text-[#535862] "
            >
              Post content
              <textarea
                value={body}
                onChange={(e) => setBody(e.target.value)}
                placeholder="Write something mind-blowing"
                className="border border-[#E2E8F0] rounded-lg p-2 mt-1 w-full h-32 font-normal leading-[21px] text-sm placeholder:text-sm placeholder:leading-[21px] placeholder:font-normal px-3"
              ></textarea>
            </label>
          </div>
          {validationError && (
            <p className="text-red-500 text-sm mt-2">{validationError}</p>
          )}
          {isError && (
            <p className="text-red-500 text-sm mt-2">Failed to add new post, please try again later</p>
          )}
        </form>
        <div className="flex justify-end mt-4 gap-2">
          <button
            disabled={isPending}
            onClick={handleCloseModal}
            className="cursor-pointer px-4 py-2 border rounded-sm text-sm border-[#E2E8F0] text-[#334155]"
          >
            Cancel
          </button>
          <button onClick={handleSubmit} className="cursor-pointer text-sm px-4 py-2 text-white rounded-sm bg-[#334155]">
            {isPending ? "Publishing..." : "Publish"}
          </button>
        </div>
      </div>
    </div>
  );
}
