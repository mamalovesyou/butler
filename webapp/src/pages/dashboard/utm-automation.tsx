import {FC, useEffect, useRef, useState} from 'react';
import { Box, Container, Typography } from "@mui/material";
import { UTMLeftMenu } from "../../components/dashboard/utm-automation/utm-left-menu";
import { CampaignListTable } from "../../components/dashboard/utm-automation/campaign-list-table";
import { useConnectors } from "../../hooks/use-connectors";
import { DashboardPageInner } from "../../components/dashboard/dashboard-page-inner";
import DashboardDrawer from "../../components/dashboard/dashboard-drawer";
import UTMDetails from '../../components/dashboard/utm-automation/utm-details';
import {listAvailableSourcesRequest} from "../../features/data-sources";
import {listWorkspaceConnectorsRequest} from "../../features/connectors";
import {useCurrentWorkspace} from "../../hooks/use-workspace";
import {useDispatch} from "react-redux";

const utmAutomation = []

for (let i = 0; i < 5; i++) {
    utmAutomation.push({
        id: i,
        name: "Campaign " + i,
        status: i % 2 === 0 ? "draft" : "live"
    })
}

const UTMAutomation: FC = () => {
    const dispatch = useDispatch();
    const { connectors } = useConnectors();
    const {workspace} = useCurrentWorkspace();

    const rootRef = useRef<HTMLDivElement | null>(null);
    const [selectedConnector, setSelectedConnector] = useState(null);
    const [utms, setUTMs] = useState<any[]>([]);
    const [drawer, setDrawer] = useState<{ isOpen: boolean; orderId?: string; }>({
        isOpen: false,
        orderId: null
    });

    useEffect(() => {
        if (workspace?.id) {
            dispatch(listAvailableSourcesRequest());
            dispatch(listWorkspaceConnectorsRequest({workspaceId: workspace.id}));
        }
    }, [workspace]);

    const handleOpenDrawer = (orderId: string): void => {
        setDrawer({
            isOpen: true,
            orderId
        });
    };

    const handleCloseDrawer = () => {
        setDrawer({
            isOpen: false,
            orderId: null
        });
    };

    return (
        <Box
            component="main"
            ref={rootRef}
            sx={{
                backgroundColor: 'background.paper',
                display: 'flex',
                flexGrow: 1,
                overflow: 'hidden',
                height: '100%',
                maxHeight: '100%'
            }}
        >
            <DashboardPageInner open={drawer.isOpen} sx={{
                height: '100%',
                maxHeight: '100%'
            }}>
                <Box sx={{
                    display: 'flex',
                    flexDirection: 'row',
                    height: '100%',
                    maxHeight: '100%',
                    p: 0,
                }}>
                    <UTMLeftMenu
                        connectors={Object.values(connectors).sort((a, b) => (a.name > b.name) ? 1 : -1)} />
                    <Container maxWidth="xl" sx={{
                        py: 4
                    }}>
                        <Typography variant="h4">Campaigns</Typography>
                        <CampaignListTable sx={{ py: 3 }} onCampaignOpen={handleOpenDrawer} campaigns={utmAutomation} />
                    </Container>
                </Box>
            </DashboardPageInner>
            <DashboardDrawer
                containerRef={rootRef}
                onClose={handleCloseDrawer}
                open={drawer.isOpen}
                title={"UTM Automation Configuration"}
            >
                <UTMDetails />
            </DashboardDrawer>
        </Box>
    );
};

export default UTMAutomation;
