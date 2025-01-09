import {
  Layout,
  Menu,
  Card,
  Row,
  Col,
  Table,
  Button,
  Modal,
  Form,
  InputNumber,
  Input,
  Space,
} from "antd";
import { Header } from "antd/es/layout/layout";
import React, { useEffect, useState } from "react";
import { GetInfo } from "../../../services/http/info";
import {
  CreateCml,
  GetCml,
  ListCmls,
  UpdateCml,
} from "../../../services/http/cml";
import { CmlsInterface } from "../../../interfaces/ICml";

const Detail = () => {
  const [selectedCml, setSelectedCml] = useState<CmlsInterface | null>(null);
  const [selectdeleteID, setselectdeleteID] = useState<Number>();
  const [isCmlModalOpen, setIsCmlModalOpen] = useState(false);
  const [isModalOpen, setIsModalOpen] = useState(false);
  const showEditModal = (cml: CmlsInterface) => {
    // console.log(info)
    setSelectedCml(cml);
    console.log(cml);
    setselectdeleteID(cml.ID);
    setIsCmlModalOpen(true);
  };

  const CML_COL = [
    {
      title: "CML number",
      dataIndex: "cml_number",
      key: "cml_number",
    },
    {
      title: "CML description",
      dataIndex: "cml_description",
      key: "cml_description",
    },
    {
      title: "Actual outside diameter",
      dataIndex: "actual_outside_diameter",
      key: "actual_outside_diameter",
    },
    {
      title: "Design thickness (mm)",
      dataIndex: "design_thickness",
      key: "desigm_thickness",
    },
    {
      title: "Structural thickness (mm)",
      dataIndex: "structural_thickness",
      key: "structural_thickness",
    },
    {
      title: "Required thickness (mm) ",
      dataIndex: "required_thickness",
      key: "required_thickness",
    },
    {
      title: "",
      key: "action",
      render: (_: any, record: CmlsInterface) => {
        return (
          <Space size="middle">
            <a>VIEW TP</a>
            <a onClick={() => showEditModal(record)}>Edit</a>
            <a>Delete</a>
          </Space>
        );
      },
    },
  ];

  const [id, setId] = useState<number | undefined>(); // กำหนดให้ id เป็น undefined
  const [line, setLine] = useState<string>(); // กำหนดให้ line เป็น undefined
  const [cmls, setCmls] = useState<CmlsInterface[]>([]);
  const [form] = Form.useForm();
  const showModal = () => {
    setIsModalOpen(true);
  };

  const handleCancel = () => {
    setIsModalOpen(false);
    form.resetFields();
  };

  const handleCmlModalCancel = () => {
    setIsCmlModalOpen(false);
    setSelectedCml(null);
  };

  const handleAdd = async () => {
    try {
      const values = await form.validateFields();

      values.info_id = id;

      await CreateCml(values);

      // เรียกข้อมูลใหม่
      getCmls();

      // ปิด modal และรีเซ็ตฟอร์ม
      setIsModalOpen(false);
      form.resetFields();

      // แสดงค่าที่ส่งไปยัง CreateCml
      console.log(values);
    } catch (error) {
      console.log("Validation Failed:", error);
    }
  };

  const handleEdit = async (values: CmlsInterface) => {
    if (selectedCml) {
      try {
        await UpdateCml(values); // ฟังก์ชันอัปเดตข้อมูลในระบบ
        getICmls(); // โหลดข้อมูลใหม่
        handleCmlModalCancel(); // ปิด Modal
      } catch (error) {
        console.error("Update failed:", error);
      }
    }
  };

  const getICmls = async () => {
    let res = await ListCmls();
    if (res) {
      setCmls(res);
    }
  };

  useEffect(() => {
    const selectedId = localStorage.getItem("selectedId");

    // แปลง selectedId จาก string เป็น number
    const idAsNumber = selectedId ? parseInt(selectedId) : undefined; // เปลี่ยนเป็น undefined แทน null

    setId(idAsNumber);
  }, []);

  useEffect(() => {
    // เรียก getInfoById เมื่อ id เปลี่ยนแปลง
    if (id !== undefined) {
      getCmls();
      getInfoById();
    }
  }, [id]); // เพิ่ม id เป็น dependency
  console.log(id);
  const getInfoById = async () => {
    if (id !== undefined) {
      // ตรวจสอบว่า id มีค่า
      let res = await GetInfo(id);
      if (res) {
        setLine(res.line_number);
      }
    }
  };
  console.log(line);

  const getCmls = async () => {
    if (id !== undefined) {
      let res = await GetCml(id);
      if (Array.isArray(res)) {
        // ตรวจสอบว่าผลลัพธ์เป็นอาร์เรย์
        setCmls(res);
      } else {
        console.error("Expected an array but got:", res);
        setCmls([]); // กำหนดเป็นอาร์เรย์ว่างหากไม่เป็นอาร์เรย์
      }
    }
  };

  console.log(cmls);

  return (
    <Layout>
      <h1 style={{ marginLeft: "50px" }}>
        {" "}
        LINE NUMBER : {line || "Loading..."}
      </h1>
      <Card
        title={
          <div
            style={{
              display: "flex",
              justifyContent: "space-between",
              alignItems: "center",
            }}
          >
            <span>CML</span>
            <Button type="primary" onClick={showModal}>
              Add Data
            </Button>
            <Modal
              title="CML"
              open={isModalOpen}
              onOk={handleAdd}
              onCancel={handleCancel}
              okText="Submit"
              cancelText="Cancel"
              width={800}
            >
              <Form form={form} layout="vertical">
                <Form.Item label="CML Number" name="Cml_number">
                  <InputNumber style={{ width: "100%" }} />
                </Form.Item>
                <Form.Item label="CML Description" name="Cml_description">
                  <Input />
                </Form.Item>
              </Form>
            </Modal>
            <Modal
              title="Edit PIPING INFORMATION"
              open={isCmlModalOpen}
              onCancel={handleCmlModalCancel}
              footer={[
                <Button key="cancel" onClick={handleCmlModalCancel}>
                  Cancel
                </Button>,
                <Button key="save" type="primary">
                  Save
                </Button>,
              ]}
              width={800}
            >{selectedCml &&(
              <Form form={form} 
              layout="vertical"
              initialValues={selectedCml}
              onFinish={(values) => handleEdit(values)}
              >
                <Form.Item label="CML Number" name="cml_number">
                  <InputNumber style={{ width: "100%" }} />
                </Form.Item>
                <Form.Item label="CML Description" name="cml_description">
                  <Input />
                </Form.Item>
              </Form>
              )}
            </Modal>
          </div>
        }
      >
        <Table
          columns={CML_COL}
          dataSource={cmls}
          bordered
          pagination={{ pageSize: 5 }}
        />
      </Card>
    </Layout>
  );
};

export default Detail;
