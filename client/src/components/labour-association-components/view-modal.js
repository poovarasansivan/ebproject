import React from "react";
import { FaTimes } from "react-icons/fa";

export default function LabourListViewModal({ isOpen, onClose, labour }) {
  if (!isOpen) return null; // Don’t render when modal is closed

  return (
    <div className="fixed inset-0 bg-black bg-opacity-40 flex items-center justify-center z-50">
      <div className="bg-white w-11/12 sm:w-2/3 md:w-1/2 lg:w-1/3 rounded-lg shadow-lg relative p-6">
        {/* Close Button */}
        <button
          className="absolute top-3 right-3 text-gray-500 hover:text-gray-700"
          onClick={onClose}
        >
          <FaTimes size={18} />
        </button>

        {/* Modal Header */}
        <h2 className="text-lg sm:text-xl font-semibold text-blue-600 mb-4 border-b pb-2">
          Labour Details
        </h2>

        {/* Modal Body */}
        {labour ? (
          <div className="grid grid-cols-1 sm:grid-cols-2 gap-3 text-sm sm:text-base text-gray-700">
            <div>
              <p className="font-semibold">Labour ID:</p>
              <p>{labour.id}</p>
            </div>
            <div>
              <p className="font-semibold">Labour Name:</p>
              <p>{labour.name}</p>
            </div>
            <div>
              <p className="font-semibold">Contact No:</p>
              <p>{labour.contact}</p>
            </div>
            <div>
              <p className="font-semibold">Manager ID:</p>
              <p>{labour.managerId}</p>
            </div>
            <div>
              <p className="font-semibold">Manager Name:</p>
              <p>{labour.managerName}</p>
            </div>
            <div>
              <p className="font-semibold">Manager Contact:</p>
              <p>{labour.managerContact}</p>
            </div>
            <div className="sm:col-span-2">
              <p className="font-semibold">Email:</p>
              <p>{labour.email}</p>
            </div>
            <div>
              <p className="font-semibold">Status:</p>
              <span
                className={`inline-block px-3 py-1 text-xs sm:text-sm font-semibold rounded-full ${
                  labour.status === "Active"
                    ? "bg-green-100 text-green-700 border border-green-300"
                    : "bg-red-100 text-red-700 border border-red-300"
                }`}
              >
                {labour.status}
              </span>
            </div>
          </div>
        ) : (
          <p className="text-gray-500 text-center py-4">No data available</p>
        )}

        {/* Modal Footer */}
        <div className="mt-6 flex justify-end">
          <button
            onClick={onClose}
            className="px-4 py-2 bg-gray-600 text-white rounded-md hover:bg-gray-800 transition-colors"
          >
            Close
          </button>
        </div>
      </div>
    </div>
  );
}
