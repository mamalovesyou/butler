import { FC, useEffect } from 'react';

// TODO: move this to utility function
export const OAuthSource = 'butler-oauth-redirect';
export type OAuthMessage = {
    source: string,
    payload: string
}

const OAuthCallback: FC = () => {
    useEffect(() => {
        if (window.opener) {
            const msg: OAuthMessage = {
                source: OAuthSource,
                payload: window.location.href,
            }
            window.opener.postMessage(msg);
            window.close();
        }
    }, []);
    return (<p>Please wait...</p>)
};

export default OAuthCallback
