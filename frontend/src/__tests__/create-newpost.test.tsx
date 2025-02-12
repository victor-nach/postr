import { render, screen } from "@testing-library/react";
import { vi } from "vitest";
import userEvent from "@testing-library/user-event";
import "@testing-library/jest-dom";
import CreateNewPostCard from '../components/create-newpost';

describe("CreateNewPostCard", () => {
  it("should render the component correctly", () => {
    render(<CreateNewPostCard onClick={vi.fn()} />);

    // Check if the 'New Post' text is displayed
    const textElement = screen.getByText(/new post/i);
    expect(textElement).toBeInTheDocument();

    // Check if the Plus Icon is displayed (could also test src)
    const iconElement = screen.getByAltText(/add a new post/i);
    expect(iconElement).toBeInTheDocument();
  });

  it("should call the onClick function when clicked", async () => {
    const onClickMock = vi.fn(); // or jest.fn() if using Jest

    render(<CreateNewPostCard onClick={onClickMock} />);

    // Use user-event to simulate a user click
    const cardElement = screen.getByRole("button"); // This should now work
    await userEvent.click(cardElement); // user-event is async, so use await

    // Check if the onClick handler was called
    expect(onClickMock).toHaveBeenCalledTimes(1);
  });
});
