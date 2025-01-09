import React, { useEffect, useState } from "react";
import {
  Card,
  Layout,
  Menu,
  Table,
  Space,
  Button,
  Form,
  InputNumber,
  Input,
  Modal,
} from "antd";
import {
  ListInfos,
  CreateInfo,
  DeleteInfo,
  UpdateInfo,
  GetInfo,
} from "../../services/http/info"; // เพิ่มฟังก์ชัน Delete และ Update
import { InfosInterface } from "../../interfaces/IInfo";
import { useNavigate } from "react-router-dom";

const { Header, Content } = Layout;

const Infomation: React.FC = () => {
  const navigate = useNavigate();

  const handleDetailClick = (id: string) => {
    localStorage.setItem("selectedId", id);
    navigate(`/detail/${id}`);
  };

  const columns = [
    {
      title: "Line Number",
      dataIndex: "line_number",
      key: "id",
    },
    {
      title: "Location",
      dataIndex: "location",
      key: "location",
    },
    {
      title: "From",
      dataIndex: "from",
      key: "from",
    },
    {
      title: "To",
      key: "To",
      dataIndex: "to",
    },
    {
      title: "Pipe Size (inch)",
      key: "PipeSize",
      dataIndex: "pipe_size",
    },
    {
      title: "Service",
      key: "Service",
      dataIndex: "service",
    },
    {
      title: "Material",
      key: "Material",
      dataIndex: "material",
    },
    {
      title: "",
      key: "action",
      render: (_: any, record: InfosInterface) => {
        return (
          <Space size="middle">
            <a onClick={() => showInfoModal(record)}>Info</a>
            <a onClick={() => handleDetailClick(`${record.ID}`)}>Detail</a>
          </Space>
        );
      },
    },
  ];

  const [infos, setInfos] = useState<InfosInterface[]>([]);
  const [isModalOpen, setIsModalOpen] = useState(false);
  const [isInfoModalOpen, setIsInfoModalOpen] = useState(false);
  const [form] = Form.useForm();
  const [selectedInfo, setSelectedInfo] = useState<InfosInterface | null>(null);
  const [selectdeleteID, setselectdeleteID] = useState<Number>();

  const showModal = () => {
    setIsModalOpen(true);
  };

  const handleCancel = () => {
    setIsModalOpen(false);
    form.resetFields();
  };

  const handleAdd = async () => {
    try {
      const values = await form.validateFields();
      await CreateInfo(values);
      console.log(values);
      getInfos();
      setIsModalOpen(false);
      form.resetFields();
    } catch (error) {
      console.log("Validation Failed:", error);
    }
  };

  const showInfoModal = (info: InfosInterface) => {
    // console.log(info)
    setSelectedInfo(info);
    console.log(info);
    setselectdeleteID(info.ID);
    setIsInfoModalOpen(true);
  };

  const handleInfoModalCancel = () => {
    setIsInfoModalOpen(false);
    setSelectedInfo(null);
  };

  const handleDelete = async () => {
    if (selectedInfo) {
      await DeleteInfo(selectdeleteID); // แทนที่ด้วย ID ที่ถูกต้อง
      getInfos();
      handleInfoModalCancel();
    }
  };

  const handleEdit = async (values: InfosInterface) => {
    if (selectedInfo) {
      try {
        await UpdateInfo(values); // ฟังก์ชันอัปเดตข้อมูลในระบบ
        getInfos(); // โหลดข้อมูลใหม่
        handleInfoModalCancel(); // ปิด Modal
        console.log("Check: ", values);
      } catch (error) {
        console.error("Update failed:", error);
      }
    }
  };

  const getInfos = async () => {
    let res = await ListInfos();
    if (res) {
      setInfos(res);
    }
  };

  useEffect(() => {
    getInfos();
  }, []);

  return (
    <Layout style={{ minHeight: "100vh" }}>
      <Content
        style={{
          padding: 24,
          margin: 0,
          minHeight: "calc(100vh - 64px)",
          display: "flex",
          justifyContent: "center",
          alignItems: "center",
        }}
      >
        <Card style={{ width: "100%", borderRadius: 8 }}>
          <div
            style={{
              display: "flex",
              justifyContent: "space-between",
              marginBottom: 16,
            }}
          >
            <h3>PIPING</h3>
            <Button type="primary" onClick={showModal}>
              Add Data
            </Button>
            <Modal
              title="PIPING INFORMATION"
              open={isModalOpen}
              onOk={handleAdd}
              onCancel={handleCancel}
              okText="Submit"
              cancelText="Cancel"
              width={800}
            >
              <Form form={form} layout="vertical">
                <Form.Item label="Line Number" name="line_number" required>
                  <Input />
                </Form.Item>
                <Form.Item label="Location" name="location" required>
                  <Input />
                </Form.Item>
                <Form.Item label="From" name="from" required>
                  <Input />
                </Form.Item>
                <Form.Item label="To" name="to" required>
                  <Input />
                </Form.Item>
                <Form.Item label="Drawing Number" name="drawing_number">
                  <Input />
                </Form.Item>
                <Form.Item label="Service" name="service">
                  <Input />
                </Form.Item>
                <Form.Item label="Material" name="material">
                  <Input />
                </Form.Item>
                <Form.Item label="In Service Date" name="in_service_date">
                  <Input type="date" />
                </Form.Item>
                <Form.Item label="Pipe Size" name="pipe_size" required>
                  <InputNumber style={{ width: "100%" }} />
                </Form.Item>
                <Form.Item label="Original Thickness" name="original_thickness">
                  <InputNumber style={{ width: "100%" }} />
                </Form.Item>
                <Form.Item label="Stress" name="stress">
                  <InputNumber style={{ width: "100%" }} />
                </Form.Item>
                <Form.Item label="Joint Efficiency" name="joint_efficiency">
                  <InputNumber style={{ width: "100%" }} step={0.01} />
                </Form.Item>
                <Form.Item label="CA" name="ca">
                  <InputNumber style={{ width: "100%" }} step={0.01} />
                </Form.Item>
                <Form.Item label="Design Life" name="design_life">
                  <InputNumber style={{ width: "100%" }} />
                </Form.Item>
                <Form.Item label="Design Pressure" name="design_pressure">
                  <InputNumber style={{ width: "100%" }} />
                </Form.Item>
                <Form.Item label="Operating Pressure" name="operating_pressure">
                  <InputNumber style={{ width: "100%" }} />
                </Form.Item>
                <Form.Item label="Design Temperature" name="design_temperature">
                  <InputNumber style={{ width: "100%" }} />
                </Form.Item>
                <Form.Item
                  label="Operating Temperature"
                  name="operating_temperature"
                >
                  <InputNumber style={{ width: "100%" }} />
                </Form.Item>
              </Form>
            </Modal>
            <Modal
              title="Edit PIPING INFORMATION"
              open={isInfoModalOpen}
              onCancel={handleInfoModalCancel}
              footer={[
                <Button key="cancel" onClick={handleInfoModalCancel}>
                  Cancel
                </Button>,
                <Button key="save" type="primary" onClick={() => form.submit()}>
                  Save
                </Button>,
              ]}
              width={800}
            >
              {selectedInfo && (
                <Form
                  layout="vertical"
                  initialValues={selectedInfo}
                  form={form}
                  onFinish={(values) =>
                    handleEdit({ ...selectedInfo, ...values })
                  }
                >
                  <Form.Item
                    label="Line Number"
                    name="line_number"
                    rules={[
                      { required: true, message: "Please enter Line Number" },
                    ]}
                  >
                    <Input />
                  </Form.Item>
                  <Form.Item
                    label="Location"
                    name="location"
                    rules={[
                      { required: true, message: "Please enter Location" },
                    ]}
                  >
                    <Input />
                  </Form.Item>
                  <Form.Item
                    label="From"
                    name="from"
                    rules={[{ required: true, message: "Please enter From" }]}
                  >
                    <Input />
                  </Form.Item>
                  <Form.Item
                    label="To"
                    name="to"
                    rules={[{ required: true, message: "Please enter To" }]}
                  >
                    <Input />
                  </Form.Item>
                  <Form.Item label="Drawing Number" name="drawing_number">
                    <Input />
                  </Form.Item>
                  <Form.Item label="Service" name="service">
                    <Input />
                  </Form.Item>
                  <Form.Item label="Material" name="material">
                    <Input />
                  </Form.Item>
                  <Form.Item label="In Service Date" name="in_service_date">
                    <Input type="date" />
                  </Form.Item>
                  <Form.Item
                    label="Pipe Size"
                    name="pipe_size"
                    rules={[
                      { required: true, message: "Please enter Pipe Size" },
                    ]}
                  >
                    <InputNumber style={{ width: "100%" }} />
                  </Form.Item>
                  <Form.Item
                    label="Original Thickness"
                    name="original_thickness"
                  >
                    <InputNumber style={{ width: "100%" }} />
                  </Form.Item>
                  <Form.Item label="Stress" name="stress">
                    <InputNumber style={{ width: "100%" }} />
                  </Form.Item>
                  <Form.Item label="Joint Efficiency" name="joint_efficiency">
                    <InputNumber style={{ width: "100%" }} step={0.01} />
                  </Form.Item>
                  <Form.Item label="CA" name="ca">
                    <InputNumber style={{ width: "100%" }} step={0.01} />
                  </Form.Item>
                  <Form.Item label="Design Life" name="design_life">
                    <InputNumber style={{ width: "100%" }} />
                  </Form.Item>
                  <Form.Item label="Design Pressure" name="design_pressure">
                    <InputNumber style={{ width: "100%" }} />
                  </Form.Item>
                  <Form.Item
                    label="Operating Pressure"
                    name="operating_pressure"
                  >
                    <InputNumber style={{ width: "100%" }} />
                  </Form.Item>
                  <Form.Item
                    label="Design Temperature"
                    name="design_temperature"
                  >
                    <InputNumber style={{ width: "100%" }} />
                  </Form.Item>
                  <Form.Item
                    label="Operating Temperature"
                    name="operating_temperature"
                  >
                    <InputNumber style={{ width: "100%" }} />
                  </Form.Item>
                </Form>
              )}
            </Modal>
          </div>
          <Table columns={columns} dataSource={infos} />
        </Card>
      </Content>
    </Layout>
  );
};

export default Infomation;
