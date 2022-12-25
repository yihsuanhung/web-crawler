import React, { useState } from "react";
import axios from "axios";

const SERVER_URL = "http://127.0.0.1:8080";

const postReq = async (url: string) => {
  const response = await axios.post(`${SERVER_URL}/crawl`, {
    url,
  });
  console.log("url", url, response);
};

export const Home = () => {
  const [url, setUrl] = useState("https://developer.mozilla.org/zh-TW/");
  return (
    <div style={{ display: "flex", gap: "12px", flexDirection: "column" }}>
      <input
        value={url}
        placeholder="輸入網址"
        style={{ width: "500px" }}
        onChange={(e) => setUrl(e.target.value)}
      />
      <div>
        <button
          style={{ margin: "6px" }}
          onClick={() => {
            postReq(url);
          }}
        >
          建立爬蟲任務
        </button>
        <button style={{ margin: "6px" }} onClick={() => setUrl("")}>
          重置網址
        </button>
      </div>
    </div>
  );
};
