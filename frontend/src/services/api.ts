import { UserProp, ApiResponseList, ApiResponse, PostProp } from "../types";


const API_KEY = import.meta.env.VITE_X_API_KEY as string;

export const fetchUsers = async (
  pageNumber: number
): Promise<{ users: UserProp[]; currentPage: number; totalPages: number }> => {
  const response = await fetch(
    `https://postr-backend-n9s0.onrender.com/users?pageNumber=${pageNumber}&pageSize=4`,
    {
      headers: {
        "X-API-Key": API_KEY, 
      },
    }
  );

  if (!response.ok) {
    throw new Error(`HTTP error! Status: ${response.status}`);
  }

  const responseData = (await response.json()) as ApiResponseList;
  return {
    users: responseData.data,
    currentPage: responseData.pagination.current_page,
    totalPages: responseData.pagination.total_pages,
  };
};

export const fetchUserPosts = async (userID: string) => {
  const response = await fetch(
    `https://postr-backend-n9s0.onrender.com/posts?${userID}`,
    {
      headers: {
        "X-API-Key": API_KEY,
      },
    }
  );

  if (!response.ok) {
    throw new Error("Network response was not ok");
  }
  const responseData = (await response.json()) as ApiResponse<PostProp[]>;
  return responseData.data;
};

export const fetchUser = async (userID: string) => {
  const response = await fetch(
    `https://postr-backend-n9s0.onrender.com/users/${userID}`,
    {
      headers: {
        "X-API-Key": API_KEY, 
      },
    }
  );

  if (!response.ok) {
    throw new Error("Network response was not ok");
  }
  const responseData = (await response.json()) as ApiResponse<UserProp>;
  return responseData.data;
};

export const deletePost = async (postID: string) => {
  const response = await fetch(
    `https://postr-backend-n9s0.onrender.com/posts/${postID}`,
    {
      method: "DELETE",
      headers: {
        "X-API-Key": API_KEY, 
      }
    }
  );

  if (!response.ok) {
    throw new Error("Failed to delete post.");
  }

  return { message: "Post deleted successfully" };
};

export const createPost = async (
  userId: string,
  title: string,
  body: string
) => {
  const response = await fetch(
    "https://postr-backend-n9s0.onrender.com/posts",
    {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        "X-API-Key": API_KEY, 
      },
      body: JSON.stringify({ userId, title, body }),
    }
  );

  if (!response.ok) {
    throw new Error("Failed to create post.");
  }

  return (await response.json()) as ApiResponse<PostProp>;
};
