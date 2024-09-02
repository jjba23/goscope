package repository

import (
	"database/sql"
	"log"

	"github.com/averageflow/goscope/v3/internal/utils"
)

// FetchDetailedRequest fetches all details from a request via its UUID.
func FetchDetailedRequest(db *sql.DB, requestUID string) DetailedRequest {
	var body string

	var headers string

	var result DetailedRequest

	row := queryDetailedRequest(db, requestUID)

	err := row.Scan(
		&result.UID,
		&result.ClientIP,
		&result.Method,
		&result.Path,
		&result.URL,
		&result.Host,
		&result.Time,
		&headers,
		&body,
		&result.Referrer,
		&result.UserAgent,
	)
	if err != nil {
		log.Println(err.Error())
	}

	result.Body = utils.PrettifyJSON(body)
	result.Headers = utils.PrettifyJSON(headers)

	return result
}

// FetchDetailedResponse fetches all details of a response via its UUID.
func FetchDetailedResponse(db *sql.DB, responseUUID string) DetailedResponse {
	var body string

	var headers string

	var result DetailedResponse

	row := queryDetailedResponse(db, responseUUID)

	err := row.Scan(
		&result.UID,
		&result.ClientIP,
		&result.Status,
		&result.Time,
		&body,
		&result.Path,
		&headers,
		&result.Size,
	)
	if err != nil {
		log.Println(err.Error())
	}

	result.Body = utils.PrettifyJSON(body)
	result.Headers = utils.PrettifyJSON(headers)

	return result
}

// FetchRequestList fetches a list of summarized requests.
func FetchRequestList(db *sql.DB, appID string, entriesPerPage, offset int) []SummarizedRequest {
	var result []SummarizedRequest

	rows, err := queryGetRequests(db, appID, entriesPerPage, offset)
	if err != nil {
		log.Println(err.Error())

		return result
	}

	if rows.Err() != nil {
		log.Println(rows.Err().Error())

		return result
	}

	defer rows.Close()

	for rows.Next() {
		var request SummarizedRequest

		err := rows.Scan(
			&request.UID,
			&request.Method,
			&request.Path,
			&request.Time,
			&request.ResponseStatus,
		)
		if err != nil {
			log.Println(err.Error())
			return result
		}

		result = append(result, request)
	}

	return result
}

// FetchSearchRequests fetches a list of summarized requests that match the input parameters of search.
func FetchSearchRequests(db *sql.DB, appID string, entriesPerPage int,
	search string, offset int, searchType int) []SummarizedRequest {
	var result []SummarizedRequest

	rows, err := querySearchRequests(db, appID, entriesPerPage, search, offset, searchType)
	if err != nil {
		log.Println(err.Error())
		return result
	}

	if rows.Err() != nil {
		log.Println(rows.Err().Error())

		return result
	}

	defer rows.Close()

	for rows.Next() {
		var request SummarizedRequest

		errr := rows.Scan(
			&request.UID,
			&request.Method,
			&request.Path,
			&request.Time,
			&request.ResponseStatus,
		)

		if errr != nil {
			log.Println(errr.Error())
		}

		result = append(result, request)
	}

	return result
}
