import { useEffect } from "react";

interface ToastProps {
  message: string;
  onClose: () => void;
  type?: "success" | "error";
  duration?: number;
}

export default function Toast({
  message,
  onClose,
  type = "success",
  duration = 1500,
}: ToastProps) {
  useEffect(() => {
    const timer = setTimeout(() => {
      onClose();
    }, duration);
    return () => clearTimeout(timer);
  }, [onClose, duration]);

  const bgColor =
    type === "success" ? "bg-green-500" : "bg-red-500";

  return (
    <div
      className={`fixed bottom-4 right-4 ${bgColor} text-white px-4 py-2 rounded-lg shadow-lg z-50 flex items-center justify-between`}
    >
      <span className="text-sm">{message}</span>
      <button
        onClick={onClose}
        className="ml-4 text-sm font-medium underline focus:outline-none"
      >
        Close
      </button>
    </div>
  );
}
