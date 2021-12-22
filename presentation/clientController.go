package presentation

import (
	"encoding/json"
	"net/http"
	"training-go-clients/entity"
	"training-go-clients/service"
	"training-go-clients/tools"
)

type ClientController struct {
	Service service.IClientService
}

func (this *ClientController) GetClient(rw http.ResponseWriter, r *http.Request) {
	id, _ := GetUIntParam(r, "id")
	client, err := this.Service.Get(id, GetExpand(r))

	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
	} else if client == nil {
		rw.WriteHeader(http.StatusNotFound)
	} else {
		str, err := json.Marshal(client)

		if err == nil {
			rw.WriteHeader(http.StatusOK)
			rw.Write(str)
		} else {
			rw.WriteHeader(http.StatusInternalServerError)
			tools.GetLogger().Println(err)
		}
	}

}

func (this *ClientController) DeleteClient(rw http.ResponseWriter, r *http.Request) {
	id, _ := GetUIntParam(r, "id")
	client, err := this.Service.Delete(id)

	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
	} else if client == nil {
		rw.WriteHeader(http.StatusNotFound)
	} else {
		str, err := json.Marshal(*client)

		if err == nil {
			rw.WriteHeader(http.StatusOK)
			rw.Write(str)
		} else {
			rw.WriteHeader(http.StatusInternalServerError)
			tools.GetLogger().Println(err)
		}
	}

}

func (this *ClientController) GetClients(rw http.ResponseWriter, r *http.Request) {
	lst, err := this.Service.GetAll(GetQuery(r))

	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
	} else {
		str, err := json.Marshal(lst)

		if err == nil {
			rw.WriteHeader(http.StatusOK)
			rw.Write(str)
		} else {
			rw.WriteHeader(http.StatusInternalServerError)
			tools.GetLogger().Println(err)
		}
	}
}

func (this *ClientController) PostClient(rw http.ResponseWriter, r *http.Request) {
	dto := &entity.Client{}
	err := json.NewDecoder(r.Body).Decode(dto)

	if err != nil {
		tools.GetLogger().Println(err)
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte(err.Error()))
		return
	}

	dto, err = this.Service.Add(dto)

	if err != nil {
		tools.GetLogger().Println(err)
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte(err.Error()))
		return
	}

	str, err := json.Marshal(*dto)

	if err == nil {
		rw.WriteHeader(http.StatusCreated)
		rw.Write(str)
	} else {
		rw.WriteHeader(http.StatusInternalServerError)
		tools.GetLogger().Println(err)
	}
}
