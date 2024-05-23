package core

import "main/db"

var client *db.PrismaClient

func GetDbClient() (*db.PrismaClient, error) {
	if client != nil {
		return client, nil
	}

	client = db.NewClient()

	err := client.Prisma.Connect()

	if err != nil {
		return &db.PrismaClient{}, err
	}

	return client, nil
}
