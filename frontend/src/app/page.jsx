"use client"
import axios from "axios";
import { useEffect, useState } from "react";

export default function Home() {

  const [datas, setdatas] = useState([])
  const fetchData = async () => {
    try {
      const res = await axios.get("http://localhost:8080/api/products");
      const data = await res.json();
      setdatas(data);
      console.log(datas);
    } catch (error) {
      console.log(error);
    }
  }

  useEffect(() => {
    fetchData()
  }, [])
  return (
   <>
   <h1>test</h1>
   </>
  );
}
