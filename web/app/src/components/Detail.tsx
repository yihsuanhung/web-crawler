import axios from "axios";
import React, { useState } from "react";
import { SERVER_URL } from "../constants";

type Data = {
  status: string;
  data: any;
};

export const Detail = () => {
  const [taskId, setTaskId] = useState("");
  const [result, setResult] = useState("");
  const [status, setStatus] = useState("");

  const sendRequest = async (id: string) => {
    try {
      const response = await axios.post<Data>(`${SERVER_URL}/status`, {
        id,
      });

      console.log("response", response);
      setResult(response.data.data);
      setStatus(response.statusText);
    } catch (error) {
      if (axios.isAxiosError(error)) {
        console.log(error.message);
        setStatus(error.message);
      }
      console.error(error);
    }
  };

  return (
    <>
      <div style={{ display: "flex", gap: "12px", alignItems: "center" }}>
        <input
          value={taskId}
          placeholder="輸入任務ID (10位字母與數字)"
          style={{ width: "300px", height: "24px" }}
          onChange={(e) => setTaskId(e.target.value)}
        />
        <div style={{ display: "flex", gap: "12px" }}>
          <button onClick={() => sendRequest(taskId)} disabled={taskId === ""}>
            查詢任務
          </button>
          <button onClick={() => setTaskId("")}>重置ID</button>
        </div>
      </div>
      {status === "" ? null : status === "OK" ? <>成功</> : <>{status}</>}
      {result !== "" && (
        <div style={{ maxWidth: "60vh", maxHeight: "40vh", overflow: "auto" }}>
          {result}
        </div>
      )}
    </>
  );
};
