import React from "react";
import { Detail, Task } from "../components";

export const Home = () => {
  return (
    <div style={{ display: "flex", flexDirection: "column", gap: "60px" }}>
      <Task />
      <Detail />
    </div>
  );
};
