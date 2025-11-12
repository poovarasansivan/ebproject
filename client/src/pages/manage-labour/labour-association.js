import React, { useState } from "react";
import { FaEdit, FaTrash, FaPlus, FaSearch, FaEye } from "react-icons/fa";
import LabourListViewModal from "../../components/labour-association-components/view-modal";
import LabourListEditModal from "../../components/labour-association-components/edit-modal";
import LabourListDeleteModal from "../../components/labour-association-components/delete-modal";
import LabourAddAssociationModal from "../../components/labour-association-components/add-labour-associtaion";

export default function LabourList() {
  const [labours] = useState([
    {
      id: "L001",
      name: "Ravi Kumar",
      contact: "9876543210",
      managerId: "M101",
      managerName: "Arun Raj",
      managerContact: "8765432109",
      email: "ravi@example.com",
      status: "Active",
    },
    {
      id: "L002",
      name: "Siva Prakash",
      contact: "9123456789",
      managerId: "M102",
      managerName: "Karthik",
      managerContact: "9988776655",
      email: "siva@example.com",
      status: "Inactive",
    },
    {
      id: "L003",
      name: "Vijay Anand",
      contact: "9871203456",
      managerId: "M103",
      managerName: "Dinesh Kumar",
      managerContact: "9876501234",
      email: "vijay@example.com",
      status: "Active",
    },
    {
      id: "L004",
      name: "Ramesh Kumar",
      contact: "7894561230",
      managerId: "M104",
      managerName: "Prakash",
      managerContact: "9876543100",
      email: "ramesh@example.com",
      status: "Active",
    },
    {
      id: "L005",
      name: "Sathish Raj",
      contact: "9800012345",
      managerId: "M105",
      managerName: "Vikram",
      managerContact: "9876509876",
      email: "sathish@example.com",
      status: "Inactive",
    },
    {
      id: "L006",
      name: "Gopi Krishnan",
      contact: "9812345678",
      managerId: "M106",
      managerName: "Arjun",
      managerContact: "9812376543",
      email: "gopi@example.com",
      status: "Active",
    },
  ]);

  const [searchTerm, setSearchTerm] = useState("");
  const [currentPage, setCurrentPage] = useState(1);
  const [selectedLabour, setSelectedLabour] = useState(null);
  const [isModalOpen, setIsModalOpen] = useState(false);
  const [isEditModalOpen, setIsEditModalOpen] = useState(false);
  const [isDeleteModalOpen, setIsDeleteModalOpen] = useState(false);
  const [isAddModalOpen, setIsAddModalOpen] = useState(false);

  const recordsPerPage = 7;

  const handleView = (labour) => {
    setSelectedLabour(labour);
    setIsModalOpen(true);
  };

  const handleCloseModal = () => {
    setSelectedLabour(null);
    setIsModalOpen(false);
  };

  const handleEdit = (labour) => {
    setSelectedLabour(labour);
    setIsEditModalOpen(true);
  };

  const handleSaveEdit = (updatedData) => {
    console.log("Edited Data:", updatedData);
    // You can later integrate API update logic here
  };

  const handleDeleteClick = (labour) => {
    setSelectedLabour(labour);
    setIsDeleteModalOpen(true);
  };

  const handleConfirmDelete = (labour) => {
    console.log("Deleted:", labour);
    // TODO: API call or local state update here
  };

  const handleSaveAssociation = (data) => {
    console.log("New Association:", data);
    // later integrate API logic here
  };

  const filteredLabours = labours.filter(
    (labour) =>
      labour.name.toLowerCase().includes(searchTerm.toLowerCase()) ||
      labour.id.toLowerCase().includes(searchTerm.toLowerCase()) ||
      labour.managerName.toLowerCase().includes(searchTerm.toLowerCase())
  );

  const indexOfLastRecord = currentPage * recordsPerPage;
  const indexOfFirstRecord = indexOfLastRecord - recordsPerPage;
  const currentLabours = filteredLabours.slice(
    indexOfFirstRecord,
    indexOfLastRecord
  );
  const totalPages = Math.ceil(filteredLabours.length / recordsPerPage);

  const handlePageChange = (page) => setCurrentPage(page);

  return (
    <>
      <div className="sm:px-4 sm:py-2 px-2 py-2 text-gray-800 font-semibold text-lg sm:text-xl">
        Labour Management
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
              className="border border-gray-300 rounded-md pl-10 pr-3 py-2 text-sm sm:text-sm w-full focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
            />
          </div>

          <button
            onClick={() => setIsAddModalOpen(true)}
            className="flex items-center justify-center gap-2 bg-blue-600 text-white text-xs sm:text-sm font-semibold px-4 py-2 rounded-md hover:bg-blue-800 transition-colors"
          >
            <FaPlus size={14} /> Add Labour
          </button>
        </div>

        <div className="overflow-x-auto">
          <table className="min-w-full text-left border border-gray-200">
            <thead className="bg-slate-200 text-gray-700 font-medium text-sm md:text-base">
              <tr>
                <th className="px-4 py-2 whitespace-nowrap">S.No</th>
                <th className="px-4 py-2 whitespace-nowrap">Labour ID</th>
                <th className="px-4 py-2 whitespace-nowrap">Labour Name</th>
                <th className="px-4 py-2 whitespace-nowrap">Contact No</th>
                <th className="px-4 py-2 whitespace-nowrap">Manager ID</th>
                <th className="px-4 py-2 whitespace-nowrap">Manager Name</th>
                <th className="px-4 py-2 whitespace-nowrap">Manager Contact</th>
                <th className="px-4 py-2 whitespace-nowrap">Email</th>
                <th className="px-4 py-2 whitespace-nowrap">Status</th>
                <th className="px-4 py-2 text-center whitespace-nowrap">
                  Actions
                </th>
              </tr>
            </thead>
            <tbody className="text-sm sm:text-md md:text-base">
              {currentLabours.length > 0 ? (
                currentLabours.map((labour, index) => (
                  <tr
                    key={labour.id}
                    className="border-b hover:bg-gray-50 transition-colors"
                  >
                    <td className="px-4 py-2">
                      {(currentPage - 1) * recordsPerPage + index + 1}
                    </td>
                    <td className="px-4 py-2 font-medium text-gray-700">
                      {labour.id}
                    </td>
                    <td className="px-4 py-2 whitespace-nowrap">
                      {labour.name}
                    </td>
                    <td className="px-4 py-2 whitespace-nowrap">
                      {labour.contact}
                    </td>
                    <td className="px-4 py-2 whitespace-nowrap">
                      {labour.managerId}
                    </td>
                    <td className="px-4 py-2 whitespace-nowrap">
                      {labour.managerName}
                    </td>
                    <td className="px-4 py-2 whitespace-nowrap">
                      {labour.managerContact}
                    </td>
                    <td className="px-4 py-2 whitespace-nowrap">
                      {labour.email}
                    </td>
                    <td className="px-4 py-2">
                      <span
                        className={`inline-block px-3 py-1 text-xs sm:text-sm font-semibold rounded-full ${
                          labour.status === "Active"
                            ? "bg-green-100 text-green-700 border border-green-300"
                            : "bg-red-100 text-red-700 border border-red-300"
                        }`}
                      >
                        {labour.status}
                      </span>
                    </td>
                    <td className="px-4 py-2 flex justify-center gap-3">
                      <button
                        className="flex items-center justify-center w-8 h-8 rounded-full bg-blue-100 text-blue-600 hover:bg-blue-200 hover:text-blue-800 transition-all duration-200"
                        title="View"
                        onClick={() => handleView(labour)}
                      >
                        <FaEye size={14} />
                      </button>
                      {/* Edit Button */}
                      <button
                        className="flex items-center justify-center w-8 h-8 rounded-full bg-blue-100 text-blue-600 hover:bg-blue-200 hover:text-blue-800 transition-all duration-200"
                        title="Edit"
                        onClick={() => handleEdit(labour)}
                      >
                        <FaEdit size={14} />
                      </button>

                      {/* Delete Button */}
                      <button
                        className="flex items-center justify-center w-8 h-8 rounded-full bg-red-100 text-red-600 hover:bg-red-200 hover:text-red-800 transition-all duration-200"
                        title="Delete"
                        onClick={() => handleDeleteClick(labour)}
                      >
                        <FaTrash size={14} />
                      </button>
                    </td>
                  </tr>
                ))
              ) : (
                <tr>
                  <td
                    colSpan="10"
                    className="text-center py-4 text-gray-500 italic"
                  >
                    No labours found.
                  </td>
                </tr>
              )}
            </tbody>
          </table>
        </div>

        <div className="flex justify-between items-center mt-4 flex-col sm:flex-row gap-2 text-sm sm:text-md md:text-base">
          {/* Showing Info */}
          <p className="text-gray-600">
            Showing{" "}
            <span className="font-medium">
              {filteredLabours.length === 0 ? 0 : indexOfFirstRecord + 1} -{" "}
              {Math.min(indexOfLastRecord, filteredLabours.length)}
            </span>{" "}
            of {filteredLabours.length}
          </p>

          {/* Pagination */}
          <div className="flex items-center gap-2">
            {/* Left Arrow */}
            <button
              onClick={() =>
                currentPage > 1 && handlePageChange(currentPage - 1)
              }
              disabled={currentPage === 1}
              className={`px-3 py-1 rounded-md border transition-all duration-200 ${
                currentPage === 1
                  ? "bg-gray-100 text-gray-400 cursor-not-allowed"
                  : "bg-white text-gray-700 hover:bg-gray-100 border-gray-300"
              }`}
            >
              ←
            </button>

            {/* Page Numbers (show 3 at a time) */}
            {Array.from({ length: totalPages }, (_, i) => i + 1)
              .slice(
                Math.max(0, currentPage - 2),
                Math.min(totalPages, currentPage + 1)
              )
              .map((page) => (
                <button
                  key={page}
                  onClick={() => handlePageChange(page)}
                  className={`px-3 py-1 rounded-md border transition-all duration-200 ${
                    page === currentPage
                      ? "bg-[#13732D] text-white border-[#13732D]"
                      : "bg-white text-gray-700 border-gray-300 hover:bg-gray-100"
                  }`}
                >
                  {page}
                </button>
              ))}

            {/* Right Arrow */}
            <button
              onClick={() =>
                currentPage < totalPages && handlePageChange(currentPage + 1)
              }
              disabled={currentPage === totalPages}
              className={`px-3 py-1 rounded-md border transition-all duration-200 ${
                currentPage === totalPages
                  ? "bg-gray-100 text-gray-400 cursor-not-allowed"
                  : "bg-white text-gray-700 hover:bg-gray-100 border-gray-300"
              }`}
            >
              →
            </button>
          </div>
        </div>
      </div>
      <LabourListViewModal
        isOpen={isModalOpen}
        onClose={handleCloseModal}
        labour={selectedLabour}
      />
      {/* Edit Modal */}
      <LabourListEditModal
        isOpen={isEditModalOpen}
        onClose={() => setIsEditModalOpen(false)}
        labour={selectedLabour}
        onSave={handleSaveEdit}
      />
      <LabourListDeleteModal
        isOpen={isDeleteModalOpen}
        onClose={() => setIsDeleteModalOpen(false)}
        labour={selectedLabour}
        onDelete={handleConfirmDelete}
      />
      <LabourAddAssociationModal
        isOpen={isAddModalOpen}
        onClose={() => setIsAddModalOpen(false)}
        onSave={handleSaveAssociation}
      />
    </>
  );
}
