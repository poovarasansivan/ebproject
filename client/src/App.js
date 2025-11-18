import "./App.css";
import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import Header from "./components/header";
import Home from "./pages/home/home";
import LabourList from "./pages/manage-labour/labour-association";
import LabourAttendance from "./pages/manage-labour/labour-attendance";
import LabourDetails from "./pages/manage-labour/labour-details";


function App() {
  return (
    <Router>
      <div className="flex-1 flex flex-col">
        <Header />
        <div className="min-h-screen bg-slate-100 py-2 px-2">
          <Routes>
            {/* <Route path="/" element={<Navigate to="/dashboard" />} /> */}
            <Route path="/home" element={<Home />} />
            <Route path="/manage-labour/labour-list" element={<LabourList />} />
            <Route path="/manage-labour/attendance" element={<LabourAttendance />} />
            <Route path="/manage-labour/details" element={<LabourDetails />}
            />
            <Route path="/" element={<Home />} />
          </Routes>
        </div>
      </div>
    </Router>
  );
}

export default App;
