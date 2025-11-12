import React, { useState } from "react";
import { FaTimes } from "react-icons/fa";

export default function LabourAddAssociationModal({ isOpen, onClose, onSave }) {
  const [formData, setFormData] = useState({
    labourid: "",
    managerid: "",
  });

  const [errors, setErrors] = useState({});

  const labourOptions = [
    { id: "L001", name: "Ravi Kumar" },
    { id: "L002", name: "Siva Prakash" },
    { id: "L003", name: "Vijay Anand" },
  ];

  const managerOptions = [
    { id: "M101", name: "Arun Raj" },
    { id: "M102", name: "Karthik" },
    { id: "M103", name: "Dinesh Kumar" },
  ];

  if (!isOpen) return null;

  const validate = () => {
    let newErrors = {};
    if (!formData.labourid) newErrors.labourid = "Please select a labour ID";
    if (!formData.managerid) newErrors.managerid = "Please select a manager ID";
    setErrors(newErrors);
    return Object.keys(newErrors).length === 0;
  };

  const handleSubmit = (e) => {
    e.preventDefault();
    if (validate()) {
      onSave(formData);
      onClose();
    }
  };

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
        <h2 className="text-lg sm:text-xl font-semibold text-blue-600 mb-4 border-b pb-2">
          Add Labour Association
        </h2>

        {/* Form */}
        <form
          onSubmit={handleSubmit}
          className="space-y-4 text-sm sm:text-base"
        >
          {/* Labour ID */}
          <div>
            <label className="font-medium text-gray-700">
              Select Labour ID <span className="text-red-500">*</span>
            </label>
            <select
              value={formData.labourid}
              onChange={(e) =>
                setFormData({ ...formData, labourid: e.target.value })
              }
              className="w-full border border-gray-300 rounded-md px-3 py-2 mt-1 focus:outline-none focus:ring-2 focus:ring-[#13732D]"
            >
              <option value="">-- Select Labour --</option>
              {labourOptions.map((labour) => (
                <option key={labour.id} value={labour.id}>
                  {labour.id} - {labour.name}
                </option>
              ))}
            </select>
            {errors.labourid && (
              <p className="text-red-500 text-xs mt-1">{errors.labourid}</p>
            )}
          </div>

          {/* Manager ID */}
          <div>
            <label className="font-medium text-gray-700">
              Select Manager ID <span className="text-red-500">*</span>
            </label>
            <select
              value={formData.managerid}
              onChange={(e) =>
                setFormData({ ...formData, managerid: e.target.value })
              }
              className="w-full border border-gray-300 rounded-md px-3 py-2 mt-1 focus:outline-none focus:ring-2 focus:ring-[#13732D]"
            >
              <option value="">-- Select Manager --</option>
              {managerOptions.map((manager) => (
                <option key={manager.id} value={manager.id}>
                  {manager.id} - {manager.name}
                </option>
              ))}
            </select>
            {errors.managerid && (
              <p className="text-red-500 text-xs mt-1">{errors.managerid}</p>
            )}
          </div>

          {/* Buttons */}
          <div className="flex justify-end gap-3 border-t pt-4 mt-4">
            <button
              type="button"
              onClick={onClose}
              className="px-4 py-2 rounded-md bg-gray-200 text-gray-700 hover:bg-gray-300 transition-colors"
            >
              Cancel
            </button>
            <button
              type="submit"
              className="px-4 py-2 rounded-md bg-[#13732D] text-white hover:bg-green-800 transition-colors"
            >
              Save
            </button>
          </div>
        </form>
      </div>
    </div>
  );
}
