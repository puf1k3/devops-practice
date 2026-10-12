import express from "express";



const app = express();



const PORT = Number(process.env.PORT || 3000);



app.get("/", (req, res) => {

  res.json({

    service: "payment-api",

    version: "1.0.0"

  });

});



app.get("/health", (req, res) => {

  res.status(200).json({ status: "healthy" });

});



app.listen(PORT, "0.0.0.0", () => {

  console.log(`Payment API running on ${PORT}`);

});
