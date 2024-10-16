"use client";
import { useEffect, useState } from "react";
import { getProducts, deleteProducts} from "./utils/api-libs"; // Pastikan ini adalah import yang benar
import Button from "@/components/Button";

export default function Home() {
  const [datas, setDatas] = useState([]); // Set default sebagai objek dengan data array

  const fetchProducts = async () => {
    try {
      const data = await getProducts(); // Ambil data dari API
      setDatas(data); // Set data yang diambil ke state
      console.log(datas);
    } catch (error) {
      console.error("Error fetching products:", error); // Tangani error
    }
  };

 const handleDelete = async (id) => {
   const result = await deleteProducts(id);
   if (result) {
     // Menghapus produk dari state setelah berhasil dihapus
     setDatas((prev) => prev.filter((product) => product.id !== id));
     console.log(`Product with ID ${id} deleted successfully.`);
   } else {
     console.log(`Failed to delete product with ID ${id}.`);
   }
 };

  useEffect(() => {
    fetchProducts(); 
  }, []);

  const addData = async () => {
    try {
      const data = await getProducts(); 
      setDatas(data);
      console.log(datas);
    } catch (error) {
      console.error("Error fetching products:", error); 
    }
  };
  

  return (
    <>
      {Array.isArray(datas.data) && datas.data.length > 0 ? (
        <div className="w-full flex flex-col justify-center items-center py-5">
          <table className="table-auto">
            <thead>
              <tr>
                <th className="px-4 py-2 border-2 bg-slate-500">ID</th>
                <th className="px-4 py-2 border-2 bg-slate-500">Name</th>
                <th className="px-4 py-2 border-2 bg-slate-500">Description</th>
                <th className="px-4 py-2 border-2 bg-slate-500">Price</th>
                <th className="px-4 py-2 border-2 bg-slate-500">Stock</th>
                <th className="px-4 py-2 border-2 bg-slate-500">Action</th>
              </tr>
            </thead>
            <tbody>
              {datas.data.map((data) => (
                <tr key={data.id}>
                  <td className="px-4 py-2 border-4">{data.id}</td>
                  <td className="px-4 py-2 border-4">{data.name}</td>
                  <td className="px-4 py-2 border-4">{data.description}</td>
                  <td className="px-4 py-2 border-4">{data.price}</td>
                  <td className="px-4 py-2 border-4">{data.stock}</td>
                  <td className="px-4 py-2 border-4 flex gap-5">
                    {" "}
                    <Button title={"update"} hover={"bg-blue-700"} />
                    <Button
                      bg="bg-red-500"
                      hover={"bg-red-700"}
                      title="Delete"
                      onClick={() => handleDelete(data.id)}
                    />
                  </td>
                </tr>
              ))}
            </tbody>
          </table>
          <Button title={"Add"} hover={"bg-blue-700"} margin="mt-5" />
        </div>
      ) : (
        <p>No products available</p> // Tampilkan pesan jika tidak ada data
      )}
    </>
  );
}
