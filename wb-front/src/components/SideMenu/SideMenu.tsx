import React, { useState } from "react";
import styles from "./sideMenu.module.scss";
import { Layout, Menu } from "antd";
import { 
  HomeOutlined, 
  FireOutlined, 
  ClockCircleOutlined, 
  MenuFoldOutlined,
  MenuUnfoldOutlined, 
  MergeOutlined
} from "@ant-design/icons";
import { useNavigate, useLocation } from "react-router-dom";

const { Sider } = Layout;

const SideMenu: React.FC = () => {
  const navigate = useNavigate();
  const location = useLocation();
  const [collapsed, setCollapsed] = useState(false);

  const items = [
    {
      key: "/home",
      icon: <HomeOutlined />,
      label: "Home",
      onClick: () => navigate("/home"),
    },
    {
      key: "/communities",
      icon: <MergeOutlined />,
      label: "Comunidades",
      onClick: () => navigate("/communities"),
    },
    {
      key: "/popular",
      icon: <FireOutlined />,
      label: "Popular",
      onClick: () => navigate("/popular"),
    },
    {
      key: "/new",
      icon: <ClockCircleOutlined />,
      label: "Novo",
      onClick: () => navigate("/new"),
    },
  ];

  return (
    <Sider 
      collapsible 
      collapsed={collapsed} 
      onCollapse={setCollapsed}
      trigger={null}
      className={styles.sideMenu}
    >
      <Menu
        mode="inline"
        items={items}
        selectedKeys={[location.pathname]}
        theme="dark"
      />
      <div 
        className={styles.trigger}
        onClick={() => setCollapsed(!collapsed)}
      >
        {collapsed ? <MenuUnfoldOutlined /> : <MenuFoldOutlined />}
      </div>
    </Sider>
  );
};

export default SideMenu;