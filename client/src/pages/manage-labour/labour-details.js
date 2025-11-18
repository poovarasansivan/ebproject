// labour details
// sno, user id, user name, contact no, email, role (admin,manager, labour), status (active/inactive), action (view, edit, delete)

// add user
// user name, contact no, email, role (dropdown - admin, manager, labour), status (dropdown - active/inactive), 

// edit user
// user name, contact no, email, role (dropdown - admin, manager, labour), status (dropdown - active/inactive)

// view user
// user id, user name, contact no, email, role (admin,manager, labour), status (active/inactive)

import React, { useState } from "react";
import { FaEdit, FaTrash, FaPlus, FaSearch, FaEye } from "react-icons/fa";

export default function LabourDetails() {
  const [users] = useState([
    {
      sno: 1,
      userId: "U001",
      userName: "Ravi Kumar",
      contact: "9876543210",
      email: "ravi@example.com",
      role: "Labour",
      status: "Active",
    },
    {
      sno: 2,
      userId: "U002",
      userName: "Siva Prakash",
      contact: "9123456789",
      email: "siva@example.com",
      role: "Manager",
      status: "Inactive",
    },
    {
      sno: 3,
      userId: "U003",
      userName: "Vijay Anand",
      contact: "9871203456",
      email: "vijay@example.com",
      role: "Admin",
      status: "Active",
    },
  ]);

  const [searchTerm, setSearchTerm] = useState("");

  const filtered = users.filter(
    (user) =>
      user.userName.toLowerCase().includes(searchTerm.toLowerCase()) ||
      user.userId.toLowerCase().includes(searchTerm.toLowerCase()) ||
      user.role.toLowerCase().includes(searchTerm.toLowerCase())
  );

  return (
    <>
      <div className="sm:px-4 sm:py-2 px-2 py-2 text-gray-800 font-semibold text-lg sm:text-xl">
        Labour Details
      </div>

      <div className="bg-white shadow-md rounded-md sm:m-3 m-2 p-4 sm:p-6">
        {/* Search & Add */}
        <div className="flex flex-col sm:flex-row sm:items-center sm:justify-between mb-4 gap-2">
          <div className="relative w-full sm:w-72">
            <FaSearch className="absolute left-3 top-3 text-gray-400" />
            <input
              type="text"
              placeholder="Search user..."
              value={searchTerm}
              onChange={(e) => setSearchTerm(e.target.value)}
              className="border border-gray-300 rounded-md pl-10 pr-3 py-2 text-sm w-full focus:outline-none focus:ring-2 focus:ring-blue-500"
            />
          </div>

          <button className="flex items-center justify-center gap-2 bg-blue-600 text-white text-xs font-semibold px-4 py-2 rounded-md hover:bg-blue-800">
            <FaPlus size={14} /> Add User
          </button>
        </div>

        {/* Table */}
        <div className="overflow-x-auto">
          <table className="min-w-full text-left border border-gray-200">
            <thead className="bg-slate-200 text-gray-700 font-medium text-sm">
              <tr>
                <th className="px-4 py-2">S.No</th>
                <th className="px-4 py-2">User ID</th>
                <th className="px-4 py-2">User Name</th>
                <th className="px-4 py-2">Contact No</th>
                <th className="px-4 py-2">Email</th>
                <th className="px-4 py-2">Role</th>
                <th className="px-4 py-2">Status</th>
                <th className="px-4 py-2 text-center">Actions</th>
              </tr>
            </thead>
            <tbody className="text-sm">
              {filtered.length > 0 ? (
                filtered.map((user) => (
                  <tr key={user.userId} className="border-b hover:bg-gray-50">
                    <td className="px-4 py-2">{user.sno}</td>
                    <td className="px-4 py-2 font-medium">{user.userId}</td>
                    <td className="px-4 py-2 whitespace-nowrap">{user.userName}</td>
                    <td className="px-4 py-2">{user.contact}</td>
                    <td className="px-4 py-2">{user.email}</td>
                    <td className="px-4 py-2">{user.role}</td>
                    <td className="px-4 py-2">
                      <span
                        className={`px-3 py-1 rounded-full text-xs font-semibold ${
                          user.status === "Active"
                            ? "bg-green-100 text-green-700 border border-green-300"
                            : "bg-red-100 text-red-700 border border-red-300"
                        }`}
                      >
                        {user.status}
                      </span>
                    </td>

                    {/* Actions */}
                    <td className="px-4 py-2 flex justify-center gap-3">
                      <button className="flex items-center justify-center w-8 h-8 rounded-full bg-blue-100 text-blue-600 hover:bg-blue-200">
                        <FaEye size={14} />
                      </button>

                      <button className="flex items-center justify-center w-8 h-8 rounded-full bg-blue-100 text-blue-600 hover:bg-blue-200">
                        <FaEdit size={14} />
                      </button>

                      <button className="flex items-center justify-center w-8 h-8 rounded-full bg-red-100 text-red-600 hover:bg-red-200">
                        <FaTrash size={14} />
                      </button>
                    </td>
                  </tr>
                ))
              ) : (
                <tr>
                  <td colSpan="8" className="text-center py-4 text-gray-500 italic">
                    No users found.
                  </td>
                </tr>
              )}
            </tbody>
          </table>
        </div>
      </div>
    </>
  );
}
