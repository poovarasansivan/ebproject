import React, { useState, useRef, useEffect } from "react";
import { FaBars, FaTimes } from "react-icons/fa";

export default function Header() {
  const [menuOpen, setMenuOpen] = useState(false);
  const [dropdownOpen, setDropdownOpen] = useState(null);
  const menuRef = useRef(null);

  const toggleMenu = () => setMenuOpen(!menuOpen);
  const toggleDropdown = (index) => {
    if (dropdownOpen === index) setDropdownOpen(null);
    else setDropdownOpen(index);
  };

  const menuItems = [
    { name: "Dashboard", link: "/" },
    {
      name: "Labour",
      subMenu: [
        { name: "Labour List", link: "/manage-labour/labour-list" },
        { name: "Attendance", link: "/manage-labour/attendance" },
        { name: "Details", link: "/manage-labour/details" },
      ],
    },
    {
      name: "Material",
      subMenu: [
        { name: "Material List", link: "/material/list" },
        { name: "Assign Material", link: "/material/assign" },
      ],
    },
    {
      name: "Projects",
      subMenu: [
        { name: "Project List", link: "/project/list" },
        { name: "Project Details", link: "/project/details" },
      ],
    },
    { name: "Expenses", link: "/expenses" },
    { name: "Reports", link: "/reports" },
  ];

  // Close dropdown when clicked outside
  useEffect(() => {
    const handleClickOutside = (event) => {
      if (menuRef.current && !menuRef.current.contains(event.target)) {
        setDropdownOpen(null);
        setMenuOpen(false);
      }
    };
    document.addEventListener("mousedown", handleClickOutside);
    return () => document.removeEventListener("mousedown", handleClickOutside);
  }, []);

  return (
    <nav
      ref={menuRef}
      className="bg-white border-b border-gray-200 relative z-50"
    >
      <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div className="flex justify-between h-16 items-center">
          {/* Logo */}
          <div className="flex-shrink-0">
            <h1 className="text-xl sm:text-2xl font-bold text-[#13732D]">
              Tenacious
            </h1>
          </div>

          {/* Desktop Menu */}
          <div className="hidden md:flex space-x-4 items-center">
            {menuItems.map((item, idx) =>
              item.subMenu ? (
                <div className="relative group" key={idx}>
                  <button
                    className="relative flex items-center space-x-1 font-semibold px-3 py-2 text-gray-700 transition-colors hover:text-blue-600"
                    onClick={() => toggleDropdown(idx)}
                  >
                    <span>{item.name}</span>
                    {/* Animated underline */}
                    <span className="absolute left-0 -bottom-0 w-0 h-[2px] bg-blue-600 transition-all duration-300 group-hover:w-full"></span>
                  </button>

                  {/* Dropdown */}
                  <div
                    className={`absolute left-0 mt-2 w-48 bg-white border rounded-md shadow-md z-50 transform transition-all duration-300 origin-top ${
                      dropdownOpen === idx
                        ? "opacity-100 scale-100 pointer-events-auto"
                        : "opacity-0 scale-95 pointer-events-none"
                    }`}
                  >
                    {item.subMenu.map((sub, sIdx) => (
                      <a
                        key={sIdx}
                        href={sub.link}
                        className="block px-4 py-2 text-gray-800 font-medium hover:bg-gray-100 transition-colors"
                      >
                        {sub.name}
                      </a>
                    ))}
                  </div>
                </div>
              ) : (
                <div className="relative group" key={idx}>
                  <a
                    href={item.link}
                    className="relative px-3 py-2 font-semibold text-gray-700 transition-colors hover:text-blue-600"
                  >
                    {item.name}
                    {/* Animated underline */}
                    <span className="absolute left-0 -bottom-0 w-0 h-[2px] bg-blue-600 transition-all duration-300 group-hover:w-full"></span>
                  </a>
                </div>
              )
            )}
          </div>

          {/* Mobile Menu Button */}
          <div className="md:hidden">
            <button
              onClick={toggleMenu}
              className="text-gray-700 focus:outline-none transition-transform duration-200"
            >
              {menuOpen ? <FaTimes size={24} className="text-red-600"/> : <FaBars size={24} className="text-blue-600"/>}
            </button>
          </div>
        </div>
      </div>

      {/* Mobile Menu */}
      <div
        className={`md:hidden bg-white border-t border-gray-200 overflow-hidden transition-all duration-300 ${
          menuOpen ? "max-h-screen opacity-100" : "max-h-0 opacity-0"
        }`}
      >
        {menuItems.map((item, idx) =>
          item.subMenu ? (
            <div key={idx} className="border-b border-gray-200">
              <button
                className="w-full text-left px-4 py-2 flex justify-between items-center text-gray-700 font-semibold"
                onClick={() => toggleDropdown(idx)}
              >
                {item.name}
                <span>{dropdownOpen === idx ? "-" : "+"}</span>
              </button>
              <div
                className={`overflow-hidden transition-all duration-300 ${
                  dropdownOpen === idx
                    ? "max-h-60 opacity-100"
                    : "max-h-0 opacity-0"
                }`}
              >
                {item.subMenu.map((sub, sIdx) => (
                  <a
                    key={sIdx}
                    href={sub.link}
                    className="block px-6 py-2 text-gray-600 hover:bg-gray-100 transition-colors"
                  >
                    {sub.name}
                  </a>
                ))}
              </div>
            </div>
          ) : (
            <a
              key={idx}
              href={item.link}
              className="block px-4 py-2 text-gray-700 border-b border-gray-200 font-semibold hover:text-blue-600 transition-colors"
            >
              {item.name}
            </a>
          )
        )}
      </div>
    </nav>
  );
}
