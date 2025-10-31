import React from 'react';
import { Layout } from 'antd';
import styles from './layout.module.scss';
import Header from '../Header/Header';
import SideMenu from '../SideMenu/SideMenu';
import { Outlet } from 'react-router-dom';

const { Content } = Layout;

const MainLayout: React.FC = () => {
  return (
    <Layout className={styles.layout}>
      <Header />
      <Layout className={styles.mainLayout}>
        <SideMenu />
        <Content className={styles.content}>
          <div className={styles.contentWrapper}>
            <Outlet />
          </div>
        </Content>
      </Layout>
    </Layout>
  );
};

export default MainLayout;