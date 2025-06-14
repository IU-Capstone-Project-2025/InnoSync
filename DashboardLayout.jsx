function DashboardLayout() {
    return (
      <div className="dashboard">
        <DashboardSidebar />
        <div className="dashboard-content">
          <Outlet />
        </div>
      </div>
    );
  }
  
  <Route element={<DashboardLayout />}>
    <Route path="/dashboard" element={<DashboardHome />} />
    <Route path="/dashboard/settings" element={<DashboardSettings />} />
  </Route>
  