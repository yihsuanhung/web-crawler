import { useState } from "react";
import axios from "axios";

const SERVER_URL = "http://127.0.0.1:8080";

export const Home = () => {
  const [url, setUrl] = useState("");
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
          onClick={() => {
            setUrl("");
            axios
              .get(`${SERVER_URL}/ping`)
              .then(function (response) {
                console.log(response);
              })
              .catch(function (error) {
                console.log(error);
              });
          }}
        >
          建立爬蟲任務
        </button>
      </div>
    </div>
  );
};
