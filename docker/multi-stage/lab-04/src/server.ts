import express from "express";

const app = express();
const PORT = 3000;

app.get("/", (req, res) => {
  res.json({ message: "User API running" });
});

app.get("/health", (req, res) => {
  res.json({ status: "ok" });
});

app.listen(PORT, "0.0.0.0", () => {
  console.log(`Server running on port ${PORT}`);
});
