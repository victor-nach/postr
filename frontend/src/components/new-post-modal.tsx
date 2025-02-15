import { useMutation, useQueryClient } from "@tanstack/react-query";
import { createPost } from "../services/api";
import { useState } from "react";
import { validatePost } from "../utils";
// import LoadingEllipsis from "../components/loadingEllipsis";

interface ModalProps {
  userId: string;
  handleCloseModal: () => void;
  setToastMessage: (message: string, type: "success" | "error") => void;
}

export default function Modal({
  userId,
  handleCloseModal,
  setToastMessage,
}: ModalProps) {
  const [title, setTitle] = useState<string>("");
  const [body, setBody] = useState<string>("");
  const [validationError, setValidationError] = useState<string>("");
  const queryClient = useQueryClient();

  const { mutate, status, isError } = useMutation({
    mutationFn: async () => await createPost(userId, title, body),

    onSuccess: async () => {
      await queryClient.invalidateQueries({ queryKey: ["userPosts", userId] });
      setToastMessage("Post created successfully!", "success");
      setTitle("");
      setBody("");
      handleCloseModal();
    },

    onError: () => {
      setToastMessage("Failed to create post. Please try again.", "error");
    },
  });

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();

    const error = validatePost(title, body);
    if (error) {
      setValidationError(error);
      return;
    }

    setValidationError("");
    mutate();
  };

  // Check if the mutation is pending
  const isPending = status === "pending";

  return (
    <div className="fixed top-0 left-0 right-0 bottom-0 flex justify-center items-center bg-[#03071282] z-50">
      <div className="bg-white p-6 rounded-lg w-[679px]">
        <h2 className="text-4xl font-medium text-[#181D27]">New Post</h2>
        <form onSubmit={handleSubmit}>
          <div className="pt-5">
            <label htmlFor="post-title" className="font-medium text-lg text-[#535862]">
              Post title
              <input
                id="post-title"
                type="text"
                name="title"
                value={title}
                placeholder="Give your post a title"
                maxLength={100}
                onChange={(e) => setTitle(e.target.value)}
                className="block border border-[#E2E8F0] rounded-[4px] pt-[10px] pr-[16px] pb-[10px] pl-[16px] mt-1 w-[631px] h-[40px] text-sm placeholder:text-sm"
              />
            </label>
          </div>
          <div className="pt-4">
            <label htmlFor="post-content" className="font-medium text-lg text-[#535862]">
              Post content
              <textarea
                id="post-content"
                value={body}
                placeholder="Write something mind-blowing"
                maxLength={500}
                onChange={(e) => setBody(e.target.value)}
                className="block border border-[#E2E8F0] rounded-[4px] pt-[10px] pr-[16px] pb-[10px] pl-[16px] mt-1 w-[631px] h-[179px] text-sm placeholder:text-sm focus:border-[#E2E8F0] focus:ring-0 outline-none resize-none"
              />
            </label>
          </div>
          {validationError && (
            <p className="text-red-500 text-sm mt-2">{validationError}</p>
          )}
          {isError && (
            <p className="text-red-500 text-sm mt-2">
              Failed to add new post. Please try again later.
            </p>
          )}
          <div className="flex justify-end mt-4 gap-2">
            <button
              type="button"
              onClick={handleCloseModal}
              disabled={isPending}
              className="px-4 py-2 border rounded-sm text-sm border-[#E2E8F0] text-[#334155]"
            >
              Cancel
            </button>
            <button
              type="submit"
              disabled={isPending}
              className={`flex items-center justify-center gap-2 px-4 py-2 text-sm rounded-sm text-white ${
                isPending ? "bg-gray-400" : "bg-[#334155]"
              }`}
            >
              <span>Publish</span>
              {isPending && (
                <span>...</span>
                // <span className="ml-2 inline-block transform scale-[0.25] origin-center">
                //   <LoadingEllipsis />
                // </span>
              )}
            </button>
          </div>
        </form>
      </div>
    </div>
  );
}
