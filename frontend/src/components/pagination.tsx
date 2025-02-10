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

  const maxPagesToShow = 10;
  const startPage =
    Math.floor((currentPage - 1) / maxPagesToShow) * maxPagesToShow + 1;
  const endPage = Math.min(startPage + maxPagesToShow - 1, totalPages);

  // Generate the page numbers to display
  const pages = Array.from(
    { length: endPage - startPage + 1 },
    (_, i) => startPage + i
  );
  console.log(totalPages, "p");

  return (
    <nav aria-label="Page navigation" className="flex w-full justify-end">
      <ul className="flex w-full justify-end">
        <li className="flex">
          <button
            onClick={handlePrev}
            className={`flex gap-1 sm:gap-3 font-semibold items-center cursor-pointer py-[10px] pe-[42px] hover:bg-[#F9F5FF]  hover:rounded-lg ${currentPage === 1 ? "opacity-50 cursor-not-allowed" : ""}`}
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

        {pages.map((pageNum) => (
          <li key={pageNum} className="flex">
            <button
              onClick={() => handleCurrent(pageNum)}
              className={`w-[40px] h-[40px] flex gap-1 sm:gap-2 font-semibold items-center justify-center cursor-pointer ${
                pageNum === currentPage ? "bg-[#F9F5FF] rounded-lg" : null
              } hover:bg-[#F9F5FF]  hover:rounded-lg`}
            >
              <span className="text-sm leading-[20px] font-semibold text-[#717680]">
                {pageNum}
              </span>
            </button>
          </li>
        ))}

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

{
  /* <div class="flex">
  <!-- Previous Button -->
  <a href="#" class="flex items-center justify-center px-3 h-8 text-sm font-medium text-gray-500 bg-white border border-gray-300 rounded-lg hover:bg-gray-100 hover:text-gray-700 dark:bg-gray-800 dark:border-gray-700 dark:text-gray-400 dark:hover:bg-gray-700 dark:hover:text-white">
    Previous
  </a>

  <!-- Next Button -->
  <a href="#" class="flex items-center justify-center px-3 h-8 ms-3 text-sm font-medium text-gray-500 bg-white border border-gray-300 rounded-lg hover:bg-gray-100 hover:text-gray-700 dark:bg-gray-800 dark:border-gray-700 dark:text-gray-400 dark:hover:bg-gray-700 dark:hover:text-white">
    Next
  </a>
</div> */
}
