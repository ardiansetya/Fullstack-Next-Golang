import { useEffect, useState } from "react";
import axios from "axios";

function App() {
  const [datas, setdatas] = useState([]);

  const fetchData = async () => {
    try {
      const response = await axios.get("http://localhost:8080/api/products"); 
      setdatas(response.data); 
    } catch (error) {
      console.error("Error fetching data:", error);
    }
  };

  useEffect(() => {
    fetchData(); 
  }, []);

  return (
      <div>
        <h1 className="text-3xl font-bold underline">Data Produk:</h1>
        <ul>
          {datas.length > 0 ? (
              datas.map((item, index) => (
                  <li key={index}>{JSON.stringify(item)}</li>
              ))
          ) : (
              <li>Loading...</li>
          )}
        </ul>
      </div>
  );
}

export default App;
