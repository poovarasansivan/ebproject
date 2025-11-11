import React from "react";
import { UserRound, Target, CircleDashed, ShieldCheck } from "lucide-react";
import Chart from "react-apexcharts";

export default function Home() {
  const cards = [
    { icon: <UserRound size={30} />, label: "Total Labours", count: 1200, className: "bg-blue-100 text-blue-600" },
    { icon: <Target size={30} />, label: "Total Projects", count: 350, className: "bg-green-100 text-green-600" },
    { icon: <CircleDashed size={30} />, label: "Pending Projects", count: 24, className: "bg-yellow-100 text-yellow-600" },
    { icon: <ShieldCheck size={30} />, label: "Active Labours", count: 58, className: "bg-red-100 text-red-600" },
  ];

  const employeeDistribution = {
    series: [
      {
        name: "Employees",
        data: [45, 32, 50, 28, 60, 40],
      },
    ],
    options: {
      chart: {
        type: "bar",
        toolbar: { show: false },
      },
      title: {
        text: "Employee Distribution Across Projects",
        align: "left",
        style: { fontSize: "14px", fontWeight: 500 },
      },
      xaxis: {
        categories: [
          "Project A",
          "Project B",
          "Project C",
          "Project D",
          "Project E",
          "Project F",
        ],
        labels: { rotate: -45 },
      },
      colors: ["#3B82F6"],
      plotOptions: {
        bar: {
          borderRadius: 6,
          columnWidth: "45%",
        },
      },
      dataLabels: { enabled: false },
      grid: { strokeDashArray: 4 },
    },
  };

  const employeeExpenses = {
    series: [
      {
        name: "Expenses (₹)",
        data: [12000, 15000, 11000, 18000, 13000, 20000, 17000],
      },
    ],
    options: {
      chart: {
        type: "line",
        toolbar: { show: false },
      },
      title: {
        text: "Employee Expenses (Last 7 Days)",
        align: "left",
        style: { fontSize: "14px", fontWeight: 500 },
      },
      xaxis: {
        categories: ["Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"],
      },
      stroke: { curve: "smooth", width: 3 },
      colors: ["#10B981"],
      markers: { size: 5 },
      grid: { strokeDashArray: 4 },
      yaxis: { labels: { formatter: (val) => `₹${val / 1000}k` } },
    },
  };

  return (
    <div className="min-h-screen bg-slate-100 py-2 px-2 lg:py-6 lg:px-6">
      <h1 className="text-lg text-gray-800 sm:text-xl font-semibold mb-6">Dashboard Overview</h1>

      <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-4 mb-6">
        {cards.map((card, index) => (
          <div
            key={index}
            className="flex items-center justify-between bg-white shadow-md rounded-2xl p-4 hover:shadow-lg transition-all"
          >
            <div className="flex items-center space-x-4">
              <div className={`p-3 rounded-full ${card.className}`}>
                {card.icon}
              </div>
              <div>
                <p className="text-gray-600 text-sm">{card.label}</p>
                <h2 className="text-md sm:text-lg font-bold">{card.count}</h2>
              </div>
            </div>
          </div>
        ))}
      </div>

      {/* Charts Section */}
      <div className="grid grid-cols-1 lg:grid-cols-2 gap-6">
        <div className="bg-white shadow-md rounded-2xl p-4">
          <Chart
            options={employeeDistribution.options}
            series={employeeDistribution.series}
            type="bar"
            height={400}
          />
        </div>

        <div className="bg-white shadow-md rounded-2xl p-4">
          <Chart
            options={employeeExpenses.options}
            series={employeeExpenses.series}
            type="line"
            height={400}
          />
        </div>
      </div>
    </div>
  );
}
