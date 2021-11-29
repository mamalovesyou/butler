import React, { useEffect, useState } from 'react';
import { OAuthSource } from '../pages/oauth-callback';

const DEFAULT_POPUP_WIDTH = 500;
const DEFAULT_POPUP_HEIGHT = 700;

type IWindowProps = {
  url: string;
  title: string;
  width?: number;
  height?: number;
};

type IPopupProps = IWindowProps & {
  onClose: () => void;
  onCode: (code: string, params: URLSearchParams) => void;
  // onError?: (error: Error) => void;
  children: React.ReactNode;
};

const createPopup = ({
  url, title, height, width,
}: IWindowProps): Window | null => {
  const left = window.screenX + (window.outerWidth - width) / 2;
  const top = window.screenY + (window.outerHeight - height) / 2.5;
  const externalPopup = window.open(
    url,
    title,
    `toolbar=no,menubar=no,width=${width},height=${height},left=${left},top=${top}`,
  );
  externalPopup.focus();
  return externalPopup;
};

export const OAuthPopup: React.FC<IPopupProps> = ({
  title = '',
  width = DEFAULT_POPUP_WIDTH,
  height = DEFAULT_POPUP_HEIGHT,
  url,
  children,
  onCode,
  onClose,
}: IPopupProps) => {
  const [externalWindow, setExternalWindow] = useState<Window | null>();
  const [messageListener, setMessageListsner] = useState<boolean>(false);

  const onContainerClick = () => {
    setExternalWindow(createPopup({
      url, title, width, height,
    }));
  };

  const receiveMessage = (event) => {
    // Do we trust the sender of this message? (might be
    // different from what we originally opened, for example).
    // console.log('event.origin', event.origin)
    // if (event.origin !== BASE_URL) {
    //   return;
    // }
    const { data } = event;
    if (data.source === OAuthSource && data.payload) {
      const params = new URL(data.payload).searchParams;
      console.log('params', params.get('code'), params.get('state'))
    }
  };

  const cleanup = () => {
    window.removeEventListener('message', receiveMessage);
    if (externalWindow) {
      externalWindow.close();
    }
    if (onClose) onClose();
  };

  useEffect(() => {
    if (externalWindow) {
      window.addEventListener('message', receiveMessage, false);
    }
    return cleanup
  }, [externalWindow])

  return (
    // eslint-disable-next-line
    <div
      onClick={() => {
        onContainerClick();
      }}
    >
      {children}
    </div>
  );
};

export default OAuthPopup;
