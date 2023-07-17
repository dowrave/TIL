using System.Collections;
using System.Collections.Generic;
using UnityEngine;

public class Enemy : MonoBehaviour
{
    // Start is called before the first frame update

    public float speed = 5;
    Vector3 dir;

    // 폭발 공장 주소
    public GameObject explosionFactory;

    void Start()
    {
        
        int randValue = Random.Range(0, 10);
        if (randValue < 3)
        {
            // 플레이어 방향
            GameObject target = GameObject.Find("Player");
            dir = target.transform.position - transform.position;
            dir.Normalize();
        }
        else
        {
            // 아래 방향
            dir = Vector3.down;
        }
    }

    // Update is called once per frame
    void Update()
    {

        transform.position += dir * speed * Time.deltaTime;
    }

    // 충돌 판정 시
    private void OnCollisionEnter(Collision other)
    {

        // 폭발 이펙트를 발생시킴
        GameObject explosion = Instantiate(explosionFactory);

        // 폭발 효과를 위치시킴
        explosion.transform.position = transform.position;

        Destroy(other.gameObject);
        Destroy(gameObject);

        // 적을 잡을 때마다 점수 표시

        // 1. Scene에서 scoreManager 객체 찾아오기
        GameObject smObject = GameObject.Find("ScoreManager");

        // 2. ScoreManager 게임 오브젝트에서 얻어온다.
        ScoreManager sm = smObject.GetComponent<ScoreManager>();

        // 3. ScoreManager 호출
        sm.SetScore(sm.GetScore() + 1);


    }



}
