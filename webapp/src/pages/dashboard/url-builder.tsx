import {FC, useRef, useState} from 'react';
import {Box, Button, Container, Divider, Typography} from '@mui/material';
import {DashboardPageInner} from "../../components/dashboard/dashboard-page-inner";
import DashboardDrawer from "../../components/dashboard/dashboard-drawer";
import URLBuilderForm from "../../components/dashboard/url-builder/url-builder-form";
import {NewSourceDialog} from "../../components/dashboard/data-sources/new-source-dialog";

const URLBuilder: FC = () => {

    const rootRef = useRef<HTMLDivElement | null>(null);
    const [drawerOpen, setDrawerOpen] = useState<boolean>(false);

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
            <DashboardPageInner open={drawerOpen} sx={{
                height: '100%',
                maxHeight: '100%'
            }}>
                <Box sx={{
                    display: 'flex',
                    flexDirection: 'row',
                    height: '100%',
                    maxHeight: '100%',
                    p: 8,
                }}>
                    <Container maxWidth="md">
                        <Box
                            sx={{
                                paddingBottom: 2,
                                alignItems: 'center',
                                display: 'flex',
                                justifyContent: 'space-between',
                            }}
                        >
                            <Typography variant="h4">Customs Links</Typography>
                            <Button variant="contained" onClick={() => setDrawerOpen(true)}>Add</Button>
                        </Box>
                        <Divider />
                    </Container>
                </Box>
            </DashboardPageInner>
            <DashboardDrawer
                containerRef={rootRef}
                onClose={() => setDrawerOpen(false)}
                open={drawerOpen}
                title={"URL Builder"}
                width={600}
            >
                <URLBuilderForm />
            </DashboardDrawer>
        </Box>
    );
};

export default URLBuilder;
