import React, { useState, useEffect } from "react";
import { FaTimes } from "react-icons/fa";

export default function LabourListEditModal({ isOpen, onClose, labour, onSave }) {
  const [formData, setFormData] = useState({
    name: "",
    contact: "",
    managerName: "",
    managerContact: "",
    managerEmail: "",
    status: "Active",
  });

  const [errors, setErrors] = useState({});

  useEffect(() => {
    if (labour) {
      setFormData({
        name: labour.name || "",
        contact: labour.contact || "",
        managerName: labour.managerName || "",
        managerContact: labour.managerContact || "",
        managerEmail: labour.email || "",
        status: labour.status || "Active",
      });
    }
  }, [labour]);

  if (!isOpen) return null;

  const validate = () => {
    let newErrors = {};

    if (!formData.name.trim()) newErrors.name = "Labour name is required";
    if (!formData.contact.trim())
      newErrors.contact = "Contact number is required";
    else if (!/^[6-9]\d{9}$/.test(formData.contact))
      newErrors.contact = "Enter a valid 10-digit phone number";

    if (!formData.managerName.trim())
      newErrors.managerName = "Manager name is required";

    if (!formData.managerContact.trim())
      newErrors.managerContact = "Manager contact is required";
    else if (!/^[6-9]\d{9}$/.test(formData.managerContact))
      newErrors.managerContact = "Enter a valid 10-digit phone number";

    if (!formData.managerEmail.trim())
      newErrors.managerEmail = "Manager email is required";
    else if (!/^[^\s@]+@[^\s@]+\.[^\s@]+$/.test(formData.managerEmail))
      newErrors.managerEmail = "Enter a valid email address";

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
      <div className="bg-white w-11/12 sm:w-2/3 md:w-1/2 lg:w-1/3 rounded-lg shadow-lg relative p-6">
        {/* Close Button */}
        <button
          className="absolute top-3 right-3 text-gray-500 hover:text-gray-700"
          onClick={onClose}
        >
          <FaTimes size={18} />
        </button>

        {/* Header */}
        <h2 className="text-lg sm:text-xl font-semibold text-blue-600 mb-4 border-b pb-2">
          Edit Labour Details
        </h2>

        {/* Form */}
        <form
          onSubmit={handleSubmit}
          className="space-y-4 text-sm sm:text-base"
        >
          {/* Labour Name */}
          <div>
            <label className="font-medium text-gray-700">
              Labour Name <span className="text-red-500">*</span>
            </label>
            <input
              type="text"
              value={formData.name}
              onChange={(e) =>
                setFormData({ ...formData, name: e.target.value })
              }
              className="w-full border border-gray-300 rounded-md px-3 py-2 mt-1 focus:outline-none focus:ring-2 focus:ring-blue-600"
            />
            {errors.name && (
              <p className="text-red-500 text-xs mt-1">{errors.name}</p>
            )}
          </div>

          {/* Contact No */}
          <div>
            <label className="font-medium text-gray-700">
              Contact No <span className="text-red-500">*</span>
            </label>
            <input
              type="text"
              value={formData.contact}
              onChange={(e) =>
                setFormData({ ...formData, contact: e.target.value })
              }
              className="w-full border border-gray-300 rounded-md px-3 py-2 mt-1 focus:outline-none focus:ring-2 focus:ring-blue-600"
            />
            {errors.contact && (
              <p className="text-red-500 text-xs mt-1">{errors.contact}</p>
            )}
          </div>

          {/* Manager Name */}
          <div>
            <label className="font-medium text-gray-700">
              Manager Name <span className="text-red-500">*</span>
            </label>
            <input
              type="text"
              value={formData.managerName}
              onChange={(e) =>
                setFormData({ ...formData, managerName: e.target.value })
              }
              className="w-full border border-gray-300 rounded-md px-3 py-2 mt-1 focus:outline-none focus:ring-2 focus:ring-blue-600"
            />
            {errors.managerName && (
              <p className="text-red-500 text-xs mt-1">{errors.managerName}</p>
            )}
          </div>

          {/* Manager Contact */}
          <div>
            <label className="font-medium text-gray-700">
              Manager Contact <span className="text-red-500">*</span>
            </label>
            <input
              type="text"
              value={formData.managerContact}
              onChange={(e) =>
                setFormData({ ...formData, managerContact: e.target.value })
              }
              className="w-full border border-gray-300 rounded-md px-3 py-2 mt-1 focus:outline-none focus:ring-2 focus:ring-blue-600"
            />
            {errors.managerContact && (
              <p className="text-red-500 text-xs mt-1">
                {errors.managerContact}
              </p>
            )}
          </div>

          {/* Manager Email */}
          <div>
            <label className="font-medium text-gray-700">
              Manager Email <span className="text-red-500">*</span>
            </label>
            <input
              type="email"
              value={formData.managerEmail}
              onChange={(e) =>
                setFormData({ ...formData, managerEmail: e.target.value })
              }
              className="w-full border border-gray-300 rounded-md px-3 py-2 mt-1 focus:outline-none focus:ring-2 focus:ring-blue-600"
            />
            {errors.managerEmail && (
              <p className="text-red-500 text-xs mt-1">{errors.managerEmail}</p>
            )}
          </div>

          {/* Status */}
          <div>
            <label className="font-medium text-gray-700">Status</label>
            <select
              value={formData.status}
              onChange={(e) =>
                setFormData({ ...formData, status: e.target.value })
              }
              className="w-full border border-gray-300 rounded-md px-3 py-2 mt-1 focus:outline-none focus:ring-2 focus:ring-blue-600"
            >
              <option value="Active">Active</option>
              <option value="Inactive">Inactive</option>
            </select>
          </div>

          {/* Buttons */}
          <div className="flex justify-end gap-3 pt-4 border-t mt-4">
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
              Save Changes
            </button>
          </div>
        </form>
      </div>
    </div>
  );
}
