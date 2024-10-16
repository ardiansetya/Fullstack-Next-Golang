export const getProducts = async () => {
   try {
      const res = await fetch("http://localhost:3001/api/products", {
         method: "GET", // Pastikan metode permintaan benar
         credentials: "include", // Sertakan cookie untuk otentikasi
      });

      if (!res.ok) {
         throw new Error(`HTTP error! status: ${res.status}`);
      }
      const data = await res.json(); // Pastikan data diset sesuai dengan struktur yang diharapkan
      console.log(data);
      return data; // Kembalikan data yang diterima
   } catch (error) {
      console.log("Error fetching data:", error);
      return null; // Mengembalikan null jika terjadi kesalahan
   }
}

export const getProductsById = async (id) => {
   try {
      const res = await fetch(`http://localhost:3001/api/products/${id}`, {
         method: "GET", // Pastikan metode permintaan benar
         credentials: "include", // Sertakan cookie untuk otentikasi
      });

      if (!res.ok) {
         throw new Error(`HTTP error! status: ${res.status}`);
      }
      const data = await res.json(); // Pastikan data diset sesuai dengan struktur yang diharapkan
      return data; // Kembalikan data yang diterima
   } catch (error) {
      console.log("Error fetching product by ID:", error);
      return null; // Mengembalikan null jika terjadi kesalahan
   }
}

export const addProducts = async (data) => {
   try {
      const res = await fetch("http://localhost:3001/api/products", {
         method: "POST", // Pastikan metode permintaan benar
         headers: {
            "Content-Type": "application/json", // Pastikan tipe konten sesuai
         },
         credentials: "include", // Sertakan cookie untuk otentikasi
         body: JSON.stringify(data), // Pastikan data diset sesuai dengan struktur yang diharapkan
      });

      if (!res.ok) {
         throw new Error(`HTTP error! status: ${res.status}`);
      }
      const result = await res.json(); // Pastikan data diset sesuai dengan struktur yang diharapkan
      return result; // Kembalikan data yang diterima
   } catch (error) {
      console.log("Error adding product:", error);
   }
}

export const deleteProducts = async (id) => {
   try {
      const res = await fetch(`http://localhost:3001/api/products/${id}`, {
         method: "DELETE", // Pastikan metode permintaan benar
         credentials: "include", // Sertakan cookie untuk otentikasi
      });

      if (!res.ok) {
         throw new Error(`HTTP error! status: ${res.status}`); // Menangani kesalahan respons
      }

      // Jika perlu, Anda dapat menambahkan logika untuk menangani hasil setelah penghapusan
      const data = await res.json(); // Jika API mengembalikan data
      console.log("Product deleted successfully:", data);
      return data; // Kembalikan data jika ada
   } catch (error) {
      console.log("Error deleting product:", error);
      return null; // Mengembalikan null jika terjadi kesalahan
   }
};
