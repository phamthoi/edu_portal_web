import { useEffect, useState } from "react";

function GradeList() {
  const [grades, setGrades] = useState([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const token = localStorage.getItem("token");

    fetch("http://localhost:8080/api/my-grades", {
      method: "GET",
      headers: {
        Authorization: `Bearer ${token}`,
      },
    })
      .then((res) => res.json())
      .then((data) => {
        setGrades(data);
        setLoading(false);
      })
      .catch((err) => {
        console.error("Lỗi khi lấy điểm:", err);
        setLoading(false);
      });
  }, []);

  if (loading) return <p className="text-center mt-10">Đang tải dữ liệu...</p>;

  if (grades.length === 0)
    return <p className="text-center mt-10">Bạn chưa có điểm học phần nào.</p>;

  return (
    <div className="max-w-2xl mx-auto mt-10">
      <h2 className="text-xl font-bold mb-4 text-center">Bảng điểm học phần</h2>
      <table className="w-full border text-sm">
        <thead>
          <tr className="bg-gray-200">
            <th className="border p-2">Học phần</th>
            <th className="border p-2">Học kỳ</th>
            <th className="border p-2">Năm học</th>
            <th className="border p-2">Điểm</th>
          </tr>
        </thead>
        <tbody>
          {grades.map((g, idx) => (
            <tr key={idx} className="text-center">
              <td className="border p-2">{g.course}</td>
              <td className="border p-2">{g.semester}</td>
              <td className="border p-2">{g.year}</td>
              <td className="border p-2">
                {g.score !== null ? g.score : "Chưa có điểm"}
              </td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
}

export default GradeList;
