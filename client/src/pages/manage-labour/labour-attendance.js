import React, { useState } from "react";
import { FaSearch } from "react-icons/fa";

export default function LabourAttendance() {
  const [records, setRecords] = useState([
    {
      sno: 1,
      labourId: "L001",
      labourName: "Ravi Kumar",
      managerId: "M101",
      managerName: "Arun Raj",
      date: "2025-11-17",
      status: "Present",
    },
    {
      sno: 2,
      labourId: "L002",
      labourName: "Siva Prakash",
      managerId: "M102",
      managerName: "Karthik",
      date: "2025-11-17",
      status: "Absent",
    },
    {
      sno: 3,
      labourId: "L003",
      labourName: "Vijay Anand",
      managerId: "M103",
      managerName: "Dinesh Kumar",
      date: "2025-11-17",
      status: "Present",
    },
  ]);

  const [searchTerm, setSearchTerm] = useState("");
  const [currentPage, setCurrentPage] = useState(1);
  const recordsPerPage = 7;

  const filtered = records.filter(
    (rec) =>
      rec.labourName.toLowerCase().includes(searchTerm.toLowerCase()) ||
      rec.labourId.toLowerCase().includes(searchTerm.toLowerCase()) ||
      rec.managerName.toLowerCase().includes(searchTerm.toLowerCase())
  );

  const indexOfLast = currentPage * recordsPerPage;
  const indexOfFirst = indexOfLast - recordsPerPage;
  const currentData = filtered.slice(indexOfFirst, indexOfLast);
  const totalPages = Math.ceil(filtered.length / recordsPerPage);

  const handlePageChange = (page) => setCurrentPage(page);

  const markStatus = (id, newStatus) => {
    setRecords((prev) =>
      prev.map((rec) =>
        rec.labourId === id ? { ...rec, status: newStatus } : rec
      )
    );
  };

  return (
    <>
      <div className="sm:px-4 sm:py-2 px-2 py-2 text-gray-800 font-semibold text-lg sm:text-xl">
        Attendance Management
      </div>

      <div className="bg-white shadow-md rounded-md sm:m-3 m-2 p-4 sm:p-6">
        <div className="flex flex-col sm:flex-row sm:items-center sm:justify-between mb-4 gap-2">
          <div className="relative w-full sm:w-72">
            <FaSearch className="absolute left-3 top-3 text-gray-400" />
            <input
              type="text"
              placeholder="Search labour or manager..."
              value={searchTerm}
              onChange={(e) => {
                setSearchTerm(e.target.value);
                setCurrentPage(1);
              }}
              className="border border-gray-300 rounded-md pl-10 pr-3 py-2 text-sm w-full focus:outline-none focus:ring-2 focus:ring-blue-500"
            />
          </div>
        </div>

        <div className="overflow-x-auto">
          <table className="min-w-full text-left border border-gray-200">
            <thead className="bg-slate-200 text-gray-700 font-medium text-sm">
              <tr>
                <th className="px-4 py-2">S.No</th>
                <th className="px-4 py-2">Labour ID</th>
                <th className="px-4 py-2">Labour Name</th>
                <th className="px-4 py-2">Manager ID</th>
                <th className="px-4 py-2">Manager Name</th>
                <th className="px-4 py-2">Date</th>
                <th className="px-4 py-2">Status</th>
                <th className="px-4 py-2 text-center">Actions</th>
              </tr>
            </thead>

            <tbody className="text-sm">
              {currentData.length > 0 ? (
                currentData.map((rec) => (
                  <tr key={rec.sno} className="border-b hover:bg-gray-50">
                    <td className="px-4 py-2">{rec.sno}</td>
                    <td className="px-4 py-2 font-medium">{rec.labourId}</td>
                    <td className="px-4 py-2 whitespace-nowrap">{rec.labourName}</td>
                    <td className="px-4 py-2">{rec.managerId}</td>
                    <td className="px-4 py-2">{rec.managerName}</td>
                    <td className="px-4 py-2">{rec.date}</td>
                    <td className="px-4 py-2">
                      <span
                        className={`px-3 py-1 rounded-full text-xs font-semibold ${
                          rec.status === "Present"
                            ? "bg-green-100 text-green-700 border border-green-300"
                            : "bg-red-100 text-red-700 border border-red-300"
                        }`}
                      >
                        {rec.status}
                      </span>
                    </td>

                    <td className="px-4 py-2 flex justify-center gap-3">
                      <button
                        onClick={() => markStatus(rec.labourId, "Present")}
                        className="flex items-center justify-center px-3 py-1 rounded-full bg-green-100 text-green-700 hover:bg-green-200 text-xs"
                      >
                        Mark Present
                      </button>
                      <button
                        onClick={() => markStatus(rec.labourId, "Absent")}
                        className="flex items-center justify-center px-3 py-1 rounded-full bg-red-100 text-red-700 hover:bg-red-200 text-xs"
                      >
                        Mark Absent
                      </button>
                    </td>
                  </tr>
                ))
              ) : (
                <tr>
                  <td colSpan="8" className="text-center py-4 text-gray-500 italic">
                    No records found.
                  </td>
                </tr>
              )}
            </tbody>
          </table>
        </div>

        <div className="flex justify-between items-center mt-4 text-sm">
          <p className="text-gray-600">
            Showing {filtered.length === 0 ? 0 : indexOfFirst + 1} -
            {Math.min(indexOfLast, filtered.length)} of {filtered.length}
          </p>

          <div className="flex items-center gap-2">
            <button
              onClick={() => currentPage > 1 && handlePageChange(currentPage - 1)}
              disabled={currentPage === 1}
              className={`px-3 py-1 rounded-md border ${
                currentPage === 1
                  ? "bg-gray-100 text-gray-400 cursor-not-allowed"
                  : "bg-white text-gray-700 hover:bg-gray-100"
              }`}
            >
              ←
            </button>

            {Array.from({ length: totalPages }, (_, i) => i + 1)
              .slice(Math.max(0, currentPage - 2), Math.min(totalPages, currentPage + 1))
              .map((page) => (
                <button
                  key={page}
                  onClick={() => handlePageChange(page)}
                  className={`px-3 py-1 rounded-md border ${
                    page === currentPage
                      ? "bg-[#13732D] text-white border-[#13732D]"
                      : "bg-white text-gray-700 hover:bg-gray-100"
                  }`}
                >
                  {page}
                </button>
              ))}

            <button
              onClick={() => currentPage < totalPages && handlePageChange(currentPage + 1)}
              disabled={currentPage === totalPages}
              className={`px-3 py-1 rounded-md border ${
                currentPage === totalPages
                  ? "bg-gray-100 text-gray-400 cursor-not-allowed"
                  : "bg-white text-gray-700 hover:bg-gray-100"
              }`}
            >
              →
            </button>
          </div>
        </div>
      </div>
    </>
  );
}