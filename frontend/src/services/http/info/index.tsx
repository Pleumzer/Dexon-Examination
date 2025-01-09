import { InfosInterface } from "../../../interfaces/IInfo";
const apiUrl = "http://localhost:8080";

async function CreateInfo(data: InfosInterface) {
    const requestOptions = {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(data),
    };
  
    let res = await fetch(`${apiUrl}/infos`, requestOptions) 
      .then((response) => response.json())
      .then((res) => {
        if (res.data) {
          return { status: true, message: res.data };
        } else {
          return { status: false, message: res.error };
        }
      });
  
    return res;
  }

async function GetInfo(id: number ) {
  const requestOptions = {
    method: "GET",
    headers: {
      "Content-Type": "application/json",
    },
  };

  let res = await fetch(`${apiUrl}/infos/${id}`, requestOptions)
    .then((response) => response.json())
    .then((res) => {
      if (res.data) {
        return res.data;
      } else {
        return false;
      }
    });

  return res;
}

async function ListInfos() {
  const requestOptions = {
    method: "GET",
    headers: {
      "Content-Type": "application/json",
    },
  };

  let res = await fetch(`${apiUrl}/infos`, requestOptions)
    .then((response) => response.json())
    .then((res) => {
      if (res.data) {
        return res.data;
      } else {
        return false;
      }
    });

  return res;
}

async function DeleteInfo(id: Number | undefined) {
  if (id === undefined) {
    return { status: false, message: "Invalid ID" };
  }

  const requestOptions = {
    method: "DELETE",
    headers: {
      "Content-Type": "application/json",
    },
  };

  try {
    const response = await fetch(`${apiUrl}/infos/${id}`, requestOptions);
    if (!response.ok) {
      // หาก status code ไม่อยู่ในช่วง 200-299
      const errorResponse = await response.json();
      return { status: false, message: errorResponse.error || "Failed to delete info" };
    }
    const data = await response.json();
    return { status: true, message: data.data || "Deleted successfully" };
  } catch (error) {
    // ใช้ type assertion เพื่อบอก TypeScript ว่า error เป็น Error
    const errorMessage = (error as Error).message || "An error occurred";
    return { status: false, message: errorMessage };
  }
}


async function UpdateInfo(data: InfosInterface ) {
  const requestOptions = {
    method: "PATCH",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify(data),
  };

  let res = await fetch(`${apiUrl}/infos`, requestOptions)
    .then((response) => response.json())
    .then((res) => {
      if (res.data) {
        return { status: true, message: res.data };
      } else {
        return { status: false, message: res.error };
      }
    });

  return res;
}

export { CreateInfo, GetInfo, ListInfos, DeleteInfo, UpdateInfo };
