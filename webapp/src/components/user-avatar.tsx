import type { FC } from "react";
import Avatar from "@mui/material/Avatar";
import PropTypes from "prop-types";
import { stringToColor, getInitials } from "../utils/string";

interface UserAvatarProps {
  name: string;
}

export const UserAvatar: FC<UserAvatarProps> = (props) => {
  const { name } = props;
  return (
    <Avatar sx={{ bgcolor: stringToColor(name) }}>{getInitials(name)}</Avatar>
  );
};

UserAvatar.propTypes = {
  name: PropTypes.string.isRequired,
};
