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

    // Check if the Plus Icon is displayed
    const iconElement = screen.getByAltText(/add a new post/i);
    expect(iconElement).toBeInTheDocument();
  });

  it("should call the onClick function when clicked", async () => {
    const onClickMock = vi.fn();

    render(<CreateNewPostCard onClick={onClickMock} />);

    // Use user-event to simulate a user click
    const cardElement = screen.getByRole("button");
    await userEvent.click(cardElement);

    expect(onClickMock).toHaveBeenCalledTimes(1);
  });
});
