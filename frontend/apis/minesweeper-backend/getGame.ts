const getGame = async (gameID: string): Promise<any> => {
  // TODO: think about replacing with config file
  const req = new Request(
    `${process.env.NEXT_PUBLIC_HOST}/api/games/${gameID}`,
    {
      method: "GET",
      headers: {
        "Content-Type": "application/json",
      },
    }
  );

  try {
    const response = await fetch(req);

    if (!response.ok || response.status !== 200) {
      return Promise.reject("Error creating new game");
    }

    // ! TODO: add game model and converter
    return response.json();
  } catch (error) {
    throw error;
  }
};

export default getGame;
