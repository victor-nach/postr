export const validatePost = (title: string, body: string) => {
  if (!title.trim()) {
    return "Title cannot be empty.";
  }
  if (!body.trim()) {
    return "Post content cannot be empty.";
  }
  return null; // No validation errors
};