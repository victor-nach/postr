import { Routes, Route } from "react-router";

import Users from "./pages/Users";
import Post from "./pages/Post";
import "./App.css";

function App() {
  return (
    <>
      <Routes>
        <Route path="/" element={<Users />} />
        <Route path="/users" element={<Users />} />
        <Route path="/users/:userID/posts" element={<Post />} />
      </Routes>
    </>
  );
}

export default App;
