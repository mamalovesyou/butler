import { FC } from 'react';
import DashboardPage from "../../components/dashboard/dashboard-page";
import ComingSoon from "../../components/dashboard/comming-soon";

const Dashboard: FC = () => {

    return (
        <DashboardPage title={"Dashboard"}><ComingSoon></ComingSoon></DashboardPage>
    );
};

export default Dashboard;
