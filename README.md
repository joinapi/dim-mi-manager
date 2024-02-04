# JOINAPI DIM MI Manager - Data In Motion Integration Manager

Manager Dashboard for DIM Integrator 
======================================

Welcome to the Manager Dashboard 4.2.0 for JOINAPI DIM MI 4.2.0.
This is a lightweight UI server that hosts a React application which is used to monitor and manage the DIM MI runtime boxes.

Configure the MI servers
======================================================================
To connect the DIM-MI boxs with the dashboard, add the following configuration to the deployment.toml file (stored in the <DIM-MI-HOME>/conf/ folder) of each server instance.

[dashboard_config]
dashboard_url = "https://{hostname/ip}:{9743}/dashboard/api/"

More information regarding the dashboard configurations can be found at

Running the Manager Dashboard
======================================================================

1. Go to the DIM_MANAGER_HOME/bin directory and run the dashboard.sh file for Linux and Unix or the dashboard.bat file for Windows.
2. Access the manager login page found at https://localhost:9743/login


--------------------------------------------------------------------------------
(c) Copyright 2024 JOINAPI.IO.