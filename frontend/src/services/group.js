async function GetGroup(groupId) {
  try {
    const response = await fetch(`/api/group/${groupId}`, {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json',
      },
    });
    const data = await response.json();
    return data;

  } catch (error) {
    nsole.error('Error fetching group:', error);
    throw error;
  }
}

async function CreateGroup(groupData) {
  try {
    const response = await fetch('/api/group', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(groupData),
    });
    const data = await response.json();
    return data;

  } catch (error) {
    console.error('Error creating group:', error);
    throw error;
  }
}

function GetAllGroups() {
  // Implement logic to fetch all groups
}

export { GetGroup, CreateGroup };
