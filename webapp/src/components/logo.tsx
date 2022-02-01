import PropTypes from "prop-types";
import {styled} from "@mui/material/styles";

interface LogoProps {
    variant?: "light" | "primary";
    height?: number;
    width?: number;
}

export const Logo = styled((props: LogoProps) => {
    const {variant, ...other} = props;

    const color = variant === "light" ? "#000000" : "#ffffff";

    return (
        <svg version="1.0" width="92.333333pt" height="110.825301pt"
             viewBox="0 0 92.333333 110.825301" preserveAspectRatio="xMidYMid meet"
             xmlns="http://www.w3.org/2000/svg"
             {...other}>
            <metadata>
                Created by potrace 1.12, written by Peter Selinger 2001-2015
            </metadata>
            <g transform="translate(0.000000,111.000000) scale(0.100000,-0.100000)" fill={color} stroke="none">
                <path
                    d="M0 721 c0 -361 1 -393 20 -452 38 -120 129 -211 249 -249 203 -64 417 46 481 249 64 203 -46 417 -249 481 -114 36 -213 22 -323 -44 l-58 -35 0 220 0 219 -60 0 -60 0 0 -389z m475 -88 c31 -10 67 -34 95 -63 106 -105 106 -265 0 -370 -105 -106 -265 -106 -370 0 -106 105 -106 265 0 371 74 74 174 97 275 62z"/>
                <path xmlns="http://www.w3.org/2000/svg"
                      d="M825 150 c-77 -30 -57 -143 25 -143 41 0 73 32 73 73 0 50 -53 88 -98 70z" fill="#10B981"/>
            </g>
        </svg>
    );
})``;

Logo.defaultProps = {
    variant: "primary",
};

Logo.propTypes = {
    variant: PropTypes.oneOf(["light", "primary"]),

};
