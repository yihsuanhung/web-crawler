import axios from "axios";
import React, { useState } from "react";
import { SERVER_URL } from "../constants";

type Data = {
  status: string;
  url: string;
  taskId: string;
};

export const Task = () => {
  const [url, setUrl] = useState("https://developer.mozilla.org/zh-TW/");
  const [status, setStatus] = useState("");
  const [taskId, setTaskId] = useState("");

  const sendRequest = async (reqUrl: string) => {
    const response = await axios.post<Data>(`${SERVER_URL}/crawl`, {
      url: reqUrl,
    });
    setStatus(response.data.status);
    setTaskId(response.data.taskId);
  };

  return (
    <div
      style={{
        display: "flex",
        gap: "24px",
        flexDirection: "column",
        alignItems: "flex-start",
      }}
    >
      <div
        style={{
          display: "flex",
          gap: "12px",
          alignItems: "center",
        }}
      >
        <input
          value={url}
          placeholder="輸入網址"
          style={{ width: "500px", height: "24px" }}
          onChange={(e) => setUrl(e.target.value)}
        />
        <div style={{ display: "flex", gap: "12px" }}>
          <button onClick={() => sendRequest(url)} disabled={url === ""}>
            建立爬蟲任務
          </button>
          <button onClick={() => setUrl("")}>重置網址</button>
        </div>
      </div>
      {taskId !== "" && status === "OK" && (
        <div>{`建立成功! 任務ID為: ${taskId}`}</div>
      )}
      {taskId !== "" && status !== "OK" && <div>建立失敗</div>}
    </div>
  );
};
