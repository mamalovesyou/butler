import { useSelector } from 'react-redux';
import { RootState } from '../features';

export const useOnboarding = () => useSelector((state: RootState) => state.onboarding);
