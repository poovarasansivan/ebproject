import React, { useState, useRef, useEffect } from "react";
import { FaBars, FaTimes, FaCogs } from "react-icons/fa";

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
    { name: "Dashboard", link: "/dashboard" },
    {
      name: "Labour",
      subMenu: [
        { name: "Labour List", link: "/labour/list" },
        { name: "Attendance", link: "/labour/attendance" },
        { name: "Details", link: "/labour/details" },
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

  // Close menu if click outside
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
    <nav ref={menuRef} className="bg-white shadow-md relative z-50">
      <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div className="flex justify-between h-16 items-center">
          {/* Logo */}
          <div className="flex-shrink-0">
            <h1 className="text-2xl font-bold text-gray-800">EB Portal</h1>
          </div>

          {/* Desktop Menu */}
          <div className="hidden md:flex space-x-4 items-center">
            {menuItems.map((item, idx) =>
              item.subMenu ? (
                <div className="relative" key={idx}>
                  <button
                    className="flex items-center space-x-1 px-3 py-2 rounded-md text-gray-700 hover:bg-gray-100 focus:outline-none"
                    onClick={() => toggleDropdown(idx)}
                  >
                    <span>{item.name}</span>
                    <FaCogs className="ml-1" />
                  </button>

                  {/* Dropdown */}
                  <div
                    className={`absolute left-0 mt-2 w-48 bg-white border rounded-md shadow-lg z-50
                      transform transition-all duration-300 origin-top
                      ${dropdownOpen === idx ? "opacity-100 scale-100 pointer-events-auto" : "opacity-0 scale-95 pointer-events-none"}`}
                  >
                    {item.subMenu.map((sub, sIdx) => (
                      <a
                        key={sIdx}
                        href={sub.link}
                        className="block px-4 py-2 text-gray-700 hover:bg-gray-100 transition-colors"
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
                  className="px-3 py-2 rounded-md text-gray-700 hover:bg-gray-100 transition-colors"
                >
                  {item.name}
                </a>
              )
            )}
          </div>

          {/* Mobile Menu Button */}
          <div className="md:hidden">
            <button
              onClick={toggleMenu}
              className="text-gray-700 focus:outline-none transition-transform duration-200"
            >
              {menuOpen ? <FaTimes size={24} /> : <FaBars size={24} />}
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
                className="w-full text-left px-4 py-2 flex justify-between items-center text-gray-700"
                onClick={() => toggleDropdown(idx)}
              >
                {item.name}
                <span>{dropdownOpen === idx ? "-" : "+"}</span>
              </button>
              <div
                className={`overflow-hidden transition-all duration-300 ${
                  dropdownOpen === idx ? "max-h-60 opacity-100" : "max-h-0 opacity-0"
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
              className="block px-4 py-2 text-gray-700 border-b border-gray-200 hover:bg-gray-100 transition-colors"
            >
              {item.name}
            </a>
          )
        )}
      </div>
    </nav>
  );
}
