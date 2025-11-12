import React from "react";
import { FaTimes } from "react-icons/fa";

export default function LabourListDeleteModal({
  isOpen,
  onClose,
  labour,
  onDelete,
}) {
  if (!isOpen) return null;

  return (
    <div className="fixed inset-0 bg-black bg-opacity-40 flex items-center justify-center z-50">
      <div className="bg-white w-11/12 sm:w-1/3 rounded-lg shadow-lg relative p-6">
        {/* Close Button */}
        <button
          className="absolute top-3 right-3 text-gray-500 hover:text-gray-700"
          onClick={onClose}
        >
          <FaTimes size={18} />
        </button>

        {/* Header */}
        <div className="flex items-center gap-2 mb-4">
          <h2 className="text-lg sm:text-xl font-semibold text-gray-800">
            Delete Confirmation
          </h2>
        </div>

        {/* Message */}
        <p className="text-gray-600 text-sm sm:text-base mb-6">
          Are you sure you want to delete{" "}
          <span className="font-semibold text-gray-800">
            {labour?.name || "this labour"}
          </span>
          ? <br />
          This action cannot be undone.
        </p>

        {/* Buttons */}
        <div className="flex justify-end gap-3 border-t pt-4">
          <button
            onClick={onClose}
            className="px-4 py-2 rounded-md bg-gray-200 text-gray-700 hover:bg-gray-300 transition-colors"
          >
            Cancel
          </button>
          <button
            onClick={() => {
              onDelete(labour);
              onClose();
            }}
            className="px-4 py-2 rounded-md bg-red-600 text-white hover:bg-red-700 transition-colors"
          >
            Delete
          </button>
        </div>
      </div>
    </div>
  );
}
