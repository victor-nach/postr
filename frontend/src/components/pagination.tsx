interface PaginationProps {
  currentPage: number;
  totalPages: number;
  setPage: (page: number) => void;
}

export default function Pagination({
  currentPage,
  totalPages,
  setPage,
}: PaginationProps) {
  const handlePrev = () => {
    if (currentPage > 1) setPage(currentPage - 1);
  };

  const handleNext = () => {
    if (currentPage < totalPages) setPage(currentPage + 1);
  };

  const handleCurrent = (pageNum: number) => setPage(pageNum);

  const generatePages = () => {
    if (totalPages <= 6) {
      return Array.from({ length: totalPages }, (_, i) => i + 1);
    }

    const pages = [];
    pages.push(1);

    if (currentPage <= 4) {
      pages.push(2, 3, 4);
      if (totalPages > 5) pages.push("...");
      pages.push(totalPages - 2, totalPages - 1, totalPages);
    } else if (currentPage >= totalPages - 3) {
      pages.push("...");
      pages.push(totalPages - 3, totalPages - 2, totalPages - 1, totalPages);
    } else {
      pages.push("...");
      pages.push(currentPage - 1, currentPage, currentPage + 1);
      pages.push("...");
      pages.push(totalPages - 1, totalPages);
    }

    return pages;
  };

  const pages = generatePages();

  return (
    <nav
      aria-label="Page navigation"
      className="flex w-full max-w-[523px] justify-end"
    >
      <ul className="flex w-full justify-end flex-wrap md:flex-nowrap">
        <li className="flex">
          <button
            onClick={handlePrev}
            className={`flex gap-1 sm:gap-3 font-semibold items-center cursor-pointer py-[10px] pe-[42px] hover:bg-[#F9F5FF]  hover:rounded-lg ${
              currentPage === 1 ? "opacity-50 cursor-not-allowed" : ""
            }`}
          >
            <svg
              width="20"
              height="20"
              viewBox="0 0 20 20"
              fill="none"
              xmlns="http://www.w3.org/2000/svg"
            >
              <path
                d="M15.8334 10H4.16669M4.16669 10L10 15.8334M4.16669 10L10 4.16669"
                stroke="#717680"
                strokeWidth="1.67"
                strokeLinecap="round"
                strokeLinejoin="round"
              />
            </svg>
            <span className="text-sm leading-[20px] font-semibold text-[#535862]">
              Previous
            </span>
          </button>
        </li>
        {pages.map((page) => {
          if (page === "...") {
            return (
              <li key="ellipsis" className="flex items-center justify-center">
                <span className="text-[#717680] px-2">…</span>
              </li>
            );
          }
          return (
            <li key={page} className="flex">
              <button
                onClick={() => handleCurrent(page as number)}
                className={`w-[40px] h-[40px] flex gap-1 sm:gap-2 font-semibold items-center justify-center cursor-pointer ${
                  page === currentPage ? "bg-[#F9F5FF] rounded-lg" : ""
                } hover:bg-[#F9F5FF] hover:rounded-lg`}
              >
                <span className="text-sm leading-[20px] font-semibold text-[#717680]">
                  {page}
                </span>
              </button>
            </li>
          );
        })}
        <li className="flex">
          <button
            onClick={handleNext}
            className={`flex sm:gap-3 font-semibold items-center cursor-pointer py-[10px] ps-[42px] hover:bg-[#F9F5FF]  hover:rounded-lg ${
              currentPage === totalPages ? "opacity-50 cursor-not-allowed" : ""
            } }`}
          >
            <span className="text-sm leading-[20px] font-semibold text-[#535862]">
              Next
            </span>
            <svg
              width="20"
              height="20"
              viewBox="0 0 20 20"
              fill="none"
              xmlns="http://www.w3.org/2000/svg"
            >
              <path
                d="M4.16663 10H15.8333M15.8333 10L9.99996 4.16669M15.8333 10L9.99996 15.8334"
                stroke="#717680"
                strokeWidth="1.67"
                strokeLinecap="round"
                strokeLinejoin="round"
              />
            </svg>
          </button>
        </li>
      </ul>
    </nav>
  );
}
